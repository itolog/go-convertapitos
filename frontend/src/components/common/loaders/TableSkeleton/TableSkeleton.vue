<script setup lang="ts">
import { Skeleton } from "@/components/ui/skeleton";
</script>

<template>
  <div class="flex flex-col h-full">
    <div class="flex py-4 gap-4">
      <Skeleton class="h-[36px] bg-muted w-full max-w-[384px] rounded-sm" />
      <Skeleton class="h-[36px] bg-muted w-full max-w-[106px] rounded-sm" />
    </div>

    <div class="flex flex-col flex-1">
      <Skeleton
        class="h-[40px] bg-muted w-full rounded-none rounded-t-sm border-b border-b-border"
      />
      <Skeleton class="h-full bg-muted w-full rounded-none rounded-b-sm border-b border-b-border" />
    </div>

    <div class="flex justify-between gap-2 py-4">
      <Skeleton class="h-[36px] bg-muted w-full rounded-sm" />
      <div class="flex space-x-2 gap-2 w-full max-w-[145px] gap">
        <Skeleton class="h-[36px] bg-muted w-full max-w-[384px] rounded-sm" />
        <Skeleton class="h-[36px] bg-muted w-full max-w-[106px] rounded-sm" />
      </div>
    </div>
  </div>
</template>

<style scoped></style>
