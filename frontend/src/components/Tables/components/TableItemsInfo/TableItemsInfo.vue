<script setup lang="ts">
import type { PropType } from "vue";

import type { ApiMeta } from "@/generated/apiClient/data-contracts.ts";

const { meta } = defineProps({
  meta: {
    type: Object as PropType<ApiMeta>,
    required: true,
  },
});
</script>

<template>
  <div
    class="flex items-center gap-1 text-sm font-bold h-9 rounded-md px-2 border text-gray-700"
  >
    Items:
    <span>{{ meta.items }}</span>
  </div>
</template>

<style scoped></style>
