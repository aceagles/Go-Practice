<script setup>
import {CheckSubmit} from "../../../wailsjs/go/main/App"
import {reactive} from 'vue'
const formData = reactive("")
</script>

<template>
    <form action="#" @submit.prevent="() => CheckSubmit(formData)">
        <label for="hostname">Hostname: </label>
        <input type="text" id="hostname" v-model="formData">
        <input type="submit" >
    </form>
</template>