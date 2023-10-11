<script setup>
import {reactive} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'
import BuySide from "./BuySide.vue";
import SellSide from "./SellSide.vue";

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}

</script>

<template>

    <el-row>
      <el-col :span="24">
        <div id="result" class="center">{{ data.resultText }}</div>
        <div id="input" class="input-box">
          <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
          <button class="btn" @click="greet">Greet</button>
        </div>
      </el-col>
    </el-row>
  <el-row>
    <el-col class="center" :span="12">
      <div >
        <BuySide/>
      </div>


    </el-col>
    <el-col class="center" :span="12">
      <div >
        <SellSide/>
      </div>

    </el-col>
  </el-row>


</template>

<style scoped>
.center {
  text-align: center;
}
</style>
