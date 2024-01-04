<script setup>
const props = defineProps(['data']);
</script>

<template>
    <div class="border m-2 p-2 bg-light">
    <table class="table">
      <thead>
        <tr>
            <th scope="col">
                Register
            </th>
          <th scope="col" v-for="(column, columnIndex) in data.RegData" :key="columnIndex">
            {{ columnIndex + data.startReg }}
          </th>
        </tr>
      </thead>
      <tbody>

        <tr>
            <th scope="row">Value</th>
          <td v-for="(column, columnIndex) in data.RegData" :key="columnIndex">
            {{ column }}
          </td>
        </tr>
      </tbody>
    </table>
    </div>
</template>