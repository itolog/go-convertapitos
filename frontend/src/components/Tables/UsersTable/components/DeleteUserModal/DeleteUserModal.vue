<script setup lang="ts">
import { UserX, Loader2 } from "lucide-vue-next";
import { type PropType, ref } from "vue";

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
import { useDeleteUser } from "@/services/api/useDleteUser.ts";
import type { User } from "@/types/user";

const isOpen = ref(false);

const { isPending, mutateAsync } = useDeleteUser();

const { user } = defineProps({
  user: {
    type: Object as PropType<User>,
    required: true,
  },
});

const userInfo = {
  ID: user?.id,
  Name: user?.name,
  Email: user?.email,
};

function stopPropagation(event: Event) {
  event.stopPropagation();
}

const emits = defineEmits(["closeDropdown"]);

async function saveAndClose() {
  if (user?.id) {
    await mutateAsync(user?.id);
    isOpen.value = false;
    emits("closeDropdown");
  }
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button
        variant="ghost"
        size="icon"
        class="flex justify-start items-center w-full gap-2 py-2"
        @click="stopPropagation"
      >
        <UserX />
        Delete User
      </Button>
    </DialogTrigger>
    <DialogContent
      class="sm:max-w-[425px] overflow-hidden grid-rows-[auto_minmax(0,1fr)_auto] p-0"
    >
      <DialogHeader class="p-6 pb-0 overflow-hidden">
        <DialogTitle>Delete profile</DialogTitle>
        <DialogDescription class="overflow-hidden">
          <ul class="flex flex-col gap-2 overflow-hidden">
            <li v-for="(value, key) in userInfo" :key="key">
              <div class="flex gap-2">
                <span class="font-bold min-w-[60px]">{{ key }}:</span>
                <span class="line-clamp-3 whitespace-normal wrap-break-word">{{
                  value
                }}</span>
              </div>
            </li>
          </ul>
        </DialogDescription>
      </DialogHeader>
      <div class="grid gap-4 py-2 overflow-y-auto px-6">
        <div class="flex flex-col font-bold text-destructive justify-between">
          <p>This action cannot be undone.</p>
          <p>The user will be permanently deleted!!!</p>
        </div>
      </div>
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
