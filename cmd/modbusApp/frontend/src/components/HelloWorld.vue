<script setup>
import {onMounted, reactive} from 'vue'
import {Greet2, GetCount} from '../../wailsjs/go/main/App'

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
function updateCounter() {
  GetCount().then(result => data.counter = result)
}

onMounted(function () {
  data.timer = setInterval(updateCounter, 500)
  console.log("Mounted")
})

</script>

<template>

    
    <p class="bg-primary">Counter: {{data.counter}}</p>
  
</template>


<style scoped>

</style>
