<script setup lang="ts">
// 接收从父组件传递的属性
import {Handle, Position} from "@vue-flow/core";

const props = defineProps({
  data: {
    label: String,
    fieldName: String,
    fieldComment: String,
    fieldType: String,
    fieldLength: String,
    required: true,
  },
  id: {
    type: String,
    required: true,
  },
});
</script>

<template>
  <div class="custom-node">
    <span>{{props.data.fieldName}}</span>
    <Handle
        type="target"
        :position="Position.Left"
        :id="id + '-target-right'"
        class="handle"
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
  position: relative;
}
</style>