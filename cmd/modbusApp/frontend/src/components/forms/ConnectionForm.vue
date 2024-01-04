<script setup>
import {CheckSubmit} from "../../../wailsjs/go/main/App"
import {reactive} from 'vue'
const formData = reactive({
    hostname: "",
    port: 502
})
</script>

<template>
    <form action="#" @submit.prevent="() => CheckSubmit(JSON.stringify(formData))">
        <label for="hostname">Hostname: </label>
        <input type="text" id="hostname" v-model="formData.hostname">
        <label for="port">Port: </label>
        <input type="number" name="port" id="port" value="502" v-model="formData.port">
        <input type="submit" >
    </form>
</template>