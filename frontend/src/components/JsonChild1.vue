<script setup lang="ts">
import { ref } from 'vue';
import {Handle, Position} from "@vue-flow/core";
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
const emit = defineEmits(['toggleChildren']);
const isExpanded = ref(true);

const toggleChildren = () => {
  isExpanded.value = !isExpanded.value;
  emit('toggleChildren', props.id, isExpanded.value);
};
</script>

<template>
  <div class="bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 shadow-sm hover:shadow-md transition-all duration-200 rounded-md p-3 relative" 
  :style="{ width: data.widthLength + 'px', height: data.heightLength + 'px' }" @click="toggleChildren">
<!--    <Fieldset :legend="props.data.JsonName" :toggleable="true" >-->
<!--      <slot></slot>-->
<!--    </Fieldset>-->
<!--<div class="flex items-center justify-between">-->
<!--      <span>{{props.data.JsonName}}</span>-->
<!--      <span>{{ isExpanded ? '−' : '+' }}</span>-->
<!--    </div>-->

    <div v-if="props.data.childrenCount > 0" class="flex items-center justify-between">
      <span>{{props.data.JsonName}}</span>
      <span>{{ isExpanded ? '−' : '+' }}</span>
    </div>
    <div v-else>
      <span>{{props.data.JsonName}}</span>
      <Handle
          type="target"
          :position="Position.Right"
          :id="id + '-target-right'"
          class="w-2.5 h-2.5 rounded-full bg-white dark:bg-gray-800 border-2 border-blue-500 hover:border-blue-600 hover:scale-110 transition-all duration-200"
      />
    </div>


  </div>
</template>

<style scoped>

</style>