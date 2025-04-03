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
  // 假设每个子节点高度为50px，上下间距为10px
  const baseHeight = 80; // 标题和边距的基础高度
  const childHeight = 30; // 每个子节点的高度（包括间距）

  console.log("字节点数量:", props.data.childrenCount)
  console.log("长度:", Math.max(96, baseHeight + props.data.childrenCount * childHeight))
  return Math.max(96, baseHeight + props.data.childrenCount * childHeight) + 'px';
});

// 根据子节点数量计算容器宽度
const containerWidth = computed(() => {
  // 基础宽度，可以根据需要调整
  const baseWidth = 220;
  return baseWidth + 'px';
});
</script>

<template>
  <div class="bg-surface-0 dark:bg-surface-900 rounded-md p-3 border border-surface-200 dark:border-surface-700" :style="{ width: containerWidth, height: containerHeight }">
    <div class="font-bold text-lg mb-2">{{props.data.tableName}}</div>
<!--    <div class="text-sm text-gray-600">{{props.data.tableComment}}</div>-->
    <Message severity="info">{{props.data.tableComment}}</Message>
  </div>
</template>

<style scoped>

</style>