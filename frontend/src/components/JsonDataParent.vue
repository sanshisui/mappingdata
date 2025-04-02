<script setup lang="ts">

// 接收从父组件传递的属性
const props = defineProps({
  data: {
    label: String,
    tableName: String,
    tableComment: String,
    required: true,
  },
  id: {
    type: String,
    required: true,
  },
});
</script>

<template>
  <div class="bg-gray-300 w-52 h-96">
    <span>{{props.data.tableName}}</span>
  </div>
</template>

<style scoped>

</style>