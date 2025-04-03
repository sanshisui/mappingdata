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
  }
});
</script>

<template>
  <div class="bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 shadow-sm hover:shadow-md transition-all duration-200 rounded-md p-3 relative" :style="{ width: data.widthLength + 'px', height: data.heightLength + 'px' }">
<!--    <Fieldset :legend="props.data.JsonName" :toggleable="true" >-->
<!--      <slot></slot>-->
<!--    </Fieldset>-->
    <span>{{props.data.JsonName}}</span>
  </div>
</template>

<style scoped>

</style>