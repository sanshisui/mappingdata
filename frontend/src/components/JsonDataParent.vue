<script setup lang="ts">
// 定义 TableData 接口
import {computed, PropType} from "vue";

interface TableData {
  label: string;
  tableName: string;
  tableComment: string;
  childrenCount: number
}


// 接收从父组件传递的属性
const props = defineProps({
  data: {
   type: Object as PropType<TableData>,
    required: true,
  },
  id: {
    type: String,
    required: true,
  },
  // 新增属性：子节点数量
  childrenCount: {
    type: Number,
    default: 0
  }
});

// 根据子节点数量计算容器高度
const containerHeight = computed(() => {
  // 基础高度 + 每个子节点的高度（包括间距）
  const baseHeight = 80; // 标题和边距的基础高度
  const childHeight = 80; // 每个子节点的高度（包括间距）

  return Math.max(120, baseHeight + props.data.childrenCount * childHeight) + 'px';
});

// 根据子节点数量计算容器宽度
const containerWidth = computed(() => {
  // 基础宽度，可以根据需要调整
  const baseWidth = 240;
  return baseWidth + 'px';
});
</script>

<template>
  <div class="bg-gray-50 dark:bg-gray-900 rounded-lg p-4 border border-gray-200 dark:border-gray-700 shadow-md" :style="{ width: containerWidth, height: containerHeight }">
    <div :style="{height: 70 + 'px' }">
      <div class="flex items-center justify-between">
        <div class="font-bold text-lg text-gray-800 dark:text-gray-100">{{props.data.tableName}}</div>
        <div v-if="props.data.childrenCount > 0" class="ml-2 px-2 py-0.5 bg-blue-100 dark:bg-blue-900 text-blue-600 dark:text-blue-300 text-xs font-medium rounded-full">
          {{props.data.childrenCount}}
        </div>
      </div>
      <div class="text-xs text-gray-500 dark:text-gray-400 mt-1 break-words" v-if="props.data.tableComment">
        {{ props.data.tableComment }}
      </div>
    </div>

  </div>
</template>