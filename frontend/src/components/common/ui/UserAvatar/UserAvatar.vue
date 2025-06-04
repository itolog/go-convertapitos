<script setup lang="ts">
import { CircleUserRound } from "lucide-vue-next";
import type { PropType } from "vue";

import { Avatar, AvatarImage } from "@/components/ui/avatar";
import type { User } from "@/types/user";

const fallbackClass = "rounded-lg w-full h-full";

const { user } = defineProps({
  user: {
    type: Object as PropType<User>,
  },
});
</script>

<template>
  <Avatar class="size-8 rounded-lg">
    <AvatarImage v-if="user?.picture" :src="user.picture" :alt="user.name" />
    <div v-else :class="fallbackClass">
      <CircleUserRound :class="fallbackClass" />
    </div>
  </Avatar>
</template>

<style scoped></style>
