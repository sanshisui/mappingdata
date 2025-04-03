<script setup lang="ts">

// 接收从父组件传递的属性
const props = defineProps({
  data: {
    type: Object,
    required: true,
  },
  id: {
    type: String,
    required: true,
  },
});
</script>

<template>
  <div class="bg-gray-50 dark:bg-gray-900 rounded-lg p-4 border border-gray-200 dark:border-gray-700 shadow-md" :style="{ width: 240 + 'px', height: 800 + 'px' }">
    <span>{{props.data.JsonName}}</span>
  </div>
</template>

<style scoped>

</style>