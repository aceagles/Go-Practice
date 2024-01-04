<script setup>
import ConnectionForm from './forms/ConnectionForm.vue'
import ModbusGroupForm from './forms/ModbusGroupForm.vue'
import RegDataList from './RegDataList.vue'
</script>

<template>
    <div>
      <connection-form />
      <modbus-group-form />
      <hr>
      <reg-data-list />
    </div>
  
</template>


<style scoped>

</style>
