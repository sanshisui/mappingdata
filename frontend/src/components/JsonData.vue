<script setup lang="ts">
import { ref } from 'vue';
import {VueFlow} from '@vue-flow/core'
import { Background } from '@vue-flow/background'
import { Controls } from '@vue-flow/controls'
import { MiniMap } from '@vue-flow/minimap'
import {Resolve} from "../../wailsjs/go/main/App";
import '@vue-flow/core/dist/style.css';
import '@vue-flow/core/dist/theme-default.css';

interface TableInfo {
  tableName: string;
  tableComment: string;
  fields: FieldInfo[];
}
interface FieldInfo {
  fieldName: string;
  fieldComment: string;
  fieldType: string;
  fieldLength: string;
}

interface NodeDefine {
  id: string;
  type: string;
  data: TableData | FieldData;
  position: Position;
  extent: string | never;
  parentNode: string | never;


}
interface TableData {
  label: string;
  tableName: string;
  tableComment: string;
}

interface FieldData {
  label: string;
  fieldName: string;
  fieldComment: string;
  fieldType: string;
  fieldLength: string;
}

interface Position {
  x: number;
  y: number;
}

const value = ref('');
const tableInfo = ref<TableInfo>();
const nodes = ref<NodeDefine>([])

function resolve() {
  Resolve(value.value).then((data:TableInfo) => {
    tableInfo.value = data;
    let tableData:TableData = {
      label: data.tableName,
      tableName: data.tableName,
      tableComment: data.tableComment,
    }
    let node:NodeDefine = {
      id: data.tableName,
      type: 'input',
      data: tableData,
      position: {x: 100, y: 100}
    }
    nodes.value.push(node);

    for (let i = 0; i < data.fields.length ; i++) {
      let fieldData:FieldData = {
        label: data.fields[i].fieldName,
        fieldComment: data.fields[i].fieldComment,
        fieldName: data.fields[i].fieldName,
        fieldLength: data.fields[i].fieldLength,
        fieldType: data.fields[i].fieldType,
      }

      nodes.value.push({
        id: data.fields[i].fieldName,
        type: 'input',
        data: fieldData,
        position: {x: 10, y: 50 + i * 50},
        extends: 'parent',
        parentNode: data.tableName
      });
    }

  })
}





</script>

<template>
  <Textarea v-model="value" rows="5" cols="30" />
  <Button label="Submit" @click="resolve" />
  <div class="simple-flow">
    <VueFlow :nodes="nodes" fit-view-on-init elevate-edges-on-select>
<!--      <template #node-parent="nodeProps">-->
<!--        <CustomNode :data="nodeProps.data" :id="nodeProps.id" />-->
<!--      </template>-->
<!--      <template #node-child="nodeProps">-->
<!--        <CustomNode1 :data="nodeProps.data" :id="nodeProps.id" />-->
<!--      </template>-->
      <MiniMap />

      <Controls />

      <Background />
    </VueFlow>
  </div>
</template>

<style scoped>
.simple-flow {
  width: 100%;
  height: 700px;
}
</style>