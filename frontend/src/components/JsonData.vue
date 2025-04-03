<script setup lang="ts">
import { ref } from 'vue';
import {CoordinateExtent, CoordinateExtentRange, VueFlow} from '@vue-flow/core'
import { Background } from '@vue-flow/background'
import { Controls } from '@vue-flow/controls'
import { MiniMap } from '@vue-flow/minimap'
import {Resolve} from "../../wailsjs/go/main/App";
import '@vue-flow/core/dist/style.css';
import '@vue-flow/core/dist/theme-default.css';
import JsonDataParent from "./JsonDataParent.vue";
import JsonDataChild from "./JsonDataChild.vue";
import JsonParent from "./JsonParent.vue";
import JsonChild1 from "./JsonChild1.vue";
import JsonChild2 from "./JsonChild2.vue";

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
  data: TableData | FieldData | JsonData;
  position: Position;
  extent?: CoordinateExtent | CoordinateExtentRange | "parent";
  parentNode?: string;
  draggable?: boolean;
  hidden?: boolean;
}

interface JsonData {
  JsonName: string;
  childrenCount: number;
  widthLength?: string;
  heightLength?: string;

}

interface TableData {
  label: string;
  tableName: string;
  tableComment: string;
  childrenCount: number;
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
const nodes = ref<NodeDefine[]>([
  {
    id: "UNB",
    type: "jsonParent",
    data: {"JsonName": "UNB", "childrenCount": 2},
    position: {x : 200, y: 200},
    draggable: true,
  },
  {
    id: "UNB01",
    type: "jsonChild1",
    data: {"JsonName": "UNB01", "childrenCount": 2, "widthLength": "210", "heightLength": "300"},
    position: {x : 10, y: 50},
    extent: 'parent',
    parentNode: 'UNB',
    draggable: false,
  },
  {
    id: "S00101",
    type: "jsonChild1",
    data: {"JsonName": "S00101", "childrenCount": 2, "widthLength": "190", "heightLength": "50"},
    position: {x : 10, y: 50},
    extent: 'parent',
    parentNode: 'UNB01',
    draggable: false,
  },
  {
    id: "S00102",
    type: "jsonChild1",
    data: {"JsonName": "S00102", "childrenCount": 2, "widthLength": "190", "heightLength": "50"},
    position: {x : 10, y: 110},
    extent: 'parent',
    parentNode: 'UNB01',
    draggable: false,
  }
]);

function resolve() {
  Resolve(value.value).then((data:TableInfo) => {
    tableInfo.value = data;
    let tableData:TableData = {
      label: data.tableName,
      tableName: data.tableName,
      tableComment: data.tableComment,
      childrenCount: data.fields.length
    }
    let node:NodeDefine = {
      id: data.tableName,
      type: 'parent',
      data: tableData,
      position: {x: 100, y: 100},
      draggable: true,
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
        type: 'child',
        data: fieldData,
        position: {x: 10, y: 80 + i * 80},
        extent: 'parent',
        parentNode: data.tableName,
        draggable: false,
      });
    }
  })
}
</script>

<template>
  <div class="bg-white dark:bg-gray-900 p-4 rounded-lg shadow-sm mb-4 border border-gray-200 dark:border-gray-700">
    <div class="flex flex-col md:flex-row gap-4 mb-4">
      <div class="flex-grow">
        <Textarea v-model="value" rows="5" class="w-full" placeholder="请输入要解析的数据..." />
      </div>
      <div class="flex items-end">
        <Button label="解析数据" icon="pi pi-search" @click="resolve" class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-md transition-colors duration-200" />
      </div>
    </div>
  </div>
  
  <div class="bg-white dark:bg-gray-900 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700">
    <div class="w-full h-[700px] rounded-lg overflow-hidden">
      <VueFlow 
        :nodes="nodes" 
        fit-view-on-init 
        elevate-edges-on-select
        :default-viewport="{ zoom: 1.5 }"
        :min-zoom="0.5"
        :max-zoom="2"
      >
        <template #node-parent="nodeProps">
          <JsonDataParent :data="nodeProps.data" :id="nodeProps.id" />
        </template>
        <template #node-child="nodeProps">
          <JsonDataChild :data="nodeProps.data" :id="nodeProps.id" />
        </template>
        <template #node-jsonParent="nodeProps">
          <JsonParent :data="nodeProps.data" :id="nodeProps.id"
          />
        </template>
        <template #node-jsonChild1="nodeProps">
          <JsonChild1 :data="nodeProps.data" :id="nodeProps.id"
          />
        </template>
<!--        <template #node-jsonChild2="nodeProps">-->
<!--          <JsonChild2 :data="nodeProps.data" :id="nodeProps.id"-->
<!--          />-->
<!--        </template>-->
        <MiniMap class="bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded shadow-sm" />
        <Controls class="bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded shadow-sm" />
        <Background />
      </VueFlow>
    </div>
  </div>
</template>