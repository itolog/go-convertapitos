<script setup lang="ts">
import type { RowSelectionState } from "@tanstack/vue-table";
import { UserX, Loader2 } from "lucide-vue-next";
import { type PropType, ref } from "vue";

import DeletionWarning from "@/components/Tables/UsersTable/components/DeletionWarning/DeletionWarning.vue";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Separator } from "@/components/ui/separator";
import { useBatchDeleteUsers } from "@/services/api/users/useBatchDeleteUsers.ts";

const { mutateAsync, isPending } = useBatchDeleteUsers();

const isOpen = ref(false);
const { users } = defineProps({
  users: {
    type: Object as PropType<RowSelectionState>,
    required: true,
  },
});

function stopPropagation(event: Event) {
  event.stopPropagation();
}

const emits = defineEmits(["deleteSuccess"]);

async function saveAndClose() {
  await mutateAsync({
    ids: Object.keys(users),
  });

  isOpen.value = false;
  emits("deleteSuccess");
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="destructive" size="icon" @click="stopPropagation">
        <UserX />
      </Button>
    </DialogTrigger>
    <DialogContent
      class="sm:max-w-[425px] overflow-hidden grid-rows-[auto_minmax(0,1fr)_auto] p-0"
    >
      <DialogHeader class="p-6 pb-0 overflow-hidden">
        <DialogTitle>Delete Profiles</DialogTitle>
        <DialogDescription class="overflow-hidden">
          <ul class="flex flex-col gap-2 overflow-hidden">
            <li v-for="value in Object.keys(users)" :key="value">
              <div class="flex gap-2">
                <span class="font-bold min-w-[20px]">ID:</span>
                <span class="line-clamp-3 whitespace-normal wrap-break-word">
                  {{ value }}
                </span>
              </div>
            </li>
          </ul>
        </DialogDescription>
      </DialogHeader>
      <div class="grid gap-4 py-2 overflow-y-auto px-6">
        <DeletionWarning />
      </div>
      <Separator />
      <DialogFooter class="p-6 pt-0">
        <Button
          :disabled="isPending"
          type="submit"
          variant="destructive"
          @click="saveAndClose"
        >
          <Loader2 v-if="isPending" class="w-4 h-4 mr-2 animate-spin" />
          Delete
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>

<style scoped></style>
