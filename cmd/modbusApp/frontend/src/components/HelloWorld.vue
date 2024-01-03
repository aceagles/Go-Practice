<script setup>
import {onMounted, reactive} from 'vue'
import {Greet2, GetCount} from '../../wailsjs/go/main/App'
import {EventsOn} from '../../wailsjs/runtime/runtime'

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
  timer: "",
  counter: 0
})

function greet() {
  Greet2(data.name).then(result => {
    data.resultText = result
  })
}

onMounted(function () {
  
  EventsOn("EmitCount", (dat) => data.counter = dat)
  console.log("Mounted")
})

</script>

<template>

    
    <p class="bg-primary">Counter: {{data.counter}}</p>
  
</template>


<style scoped>

</style>
