<script setup>
import {onMounted, ref} from 'vue'
import {UpdateGroups, StopReading, GetGroups} from "../../../wailsjs/go/main/App"
import {EventsOn} from "../../../wailsjs/runtime/runtime"
const data = ref([])

function addGroup () {
    console.log("Adding Group")
    data.value.push({startReg: 0, numReg: 0})
}

onMounted(() => {
    GetGroups().then(result => data.value = result)
})

</script>

<template>
    <div>
        <div v-for="group in data" :key="group">
            <input type="number" name="start_reg" id="reg2" v-model="group.startReg">
            <input type="number" name="num_reg" id="reg3" v-model="group.numReg">

        </div>
        <input type="button" value="Add Group" @click="addGroup">
        <input type="button" value="Update" @click="() => UpdateGroups(JSON.stringify(data))">
        <input type="button" value="Stop" @click="StopReading">
    </div>
</template>
