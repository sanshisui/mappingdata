<script lang="ts" setup>
//import HelloWorld from './components/HelloWorld.vue'
import AppTopBar from "./components/AppTopBar.vue";
import JsonData from "./components/JsonData.vue";
</script>

<template>
<!--  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>-->
  <div class="bg-surface-50 dark:bg-surface-950 min-h-screen p-8 flex flex-col gap-6">
    <AppTopBar />
    <JsonData />
  </div>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
