<script setup>
import {onMounted, reactive} from 'vue'
import {EventsOn} from '../../wailsjs/runtime/runtime'
import ConnectionForm from './forms/ConnectionForm.vue'

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
  timer: "",
  counter: 0
})


onMounted(function () {
  console.log("Mounted")
})

</script>

<template>
    <div>
      <connection-form />
    </div>
  
</template>


<style scoped>

</style>
