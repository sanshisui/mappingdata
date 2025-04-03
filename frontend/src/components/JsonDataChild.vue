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
  <div class="bg-surface-100 dark:bg-surface-700 border border-surface-300 dark:border-surface-600 rounded-md p-2.5 w-50 h-16 relative mt-12">
    <div class="flex justify-between items-center mb-1">
      <div class="font-bold text-gray-800 text-sm">{{props.data.fieldName}}</div>
      <div class="text-xs text-gray-500 bg-gray-100 px-1.5 py-0.5 rounded">{{ fieldTypeDisplay() }}</div>
    </div>
    <div class="text-xs text-gray-400 italic mt-1 break-words" v-if="props.data.fieldComment">
      {{ props.data.fieldComment }}
    </div>

    <Handle
        type="target"
        :position="Position.Left"
        :id="id + '-target-right'"
        class="w-2.5 h-2.5 rounded-full bg-white border-2 border-gray-600 transition-all duration-200"
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