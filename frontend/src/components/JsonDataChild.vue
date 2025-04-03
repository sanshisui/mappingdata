<script setup lang="ts">
// 接收从父组件传递的属性
import {Handle, Position} from "@vue-flow/core";
import {PropType} from "vue";

// 定义 FieldData 接口
interface FieldData {
  label: string;
  fieldName: string;
  fieldComment: string;
  fieldType: string;
  fieldLength: string;
}
const props = defineProps({
  data: {
    type: Object as PropType<FieldData>,
    required: true,
  },
  id: {
    type: String,
    required: true,
  },
});

// 计算完整的字段类型显示
const fieldTypeDisplay = () => {
  if (props.data.fieldLength) {
    return `${props.data.fieldType}(${props.data.fieldLength})`;
  }
  return props.data.fieldType;
};
</script>

<template>
  <div class="bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 shadow-sm hover:shadow-md transition-all duration-200 rounded-md p-3 relative" :style="{ width: 210 + 'px', height: 70 + 'px' }">
    <div class="flex justify-between items-center mb-2">
      <div class="font-semibold text-gray-800 dark:text-gray-100 text-sm">{{props.data.fieldName}}</div>
      <div class="text-xs bg-blue-50 dark:bg-blue-900 text-blue-600 dark:text-blue-300 px-2 py-0.5 rounded-full">{{ fieldTypeDisplay() }}</div>
    </div>
    <div class="text-xs text-gray-500 dark:text-gray-400 mt-1 break-words" v-if="props.data.fieldComment">
      {{ props.data.fieldComment }}
    </div>

    <Handle
        type="target"
        :position="Position.Left"
        :id="id + '-target-right'"
        class="w-2.5 h-2.5 rounded-full bg-white dark:bg-gray-800 border-2 border-blue-500 hover:border-blue-600 hover:scale-110 transition-all duration-200"
    />
  </div>
</template>

<style scoped>
/* 连接点样式 */
.handle {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background-color: #fff;
  border: 2px solid #555;
  transition: all 0.2s;
}
.custom-node {
  background: white;
  border: 1px solid #ddd;
  border-radius: 5px;
  padding: 10px;
  width: 200px;
  height: 80px;
  position: relative;
}
</style>