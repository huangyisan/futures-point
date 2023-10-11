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
        <el-button type="primary">Primary</el-button>
        <div id="result" class="result">{{ data.resultText }}</div>
        <div id="input" class="input-box">
          <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
          <button class="btn" @click="greet">Greet</button>
        </div>
      </el-col>
    </el-row>
  <el-row>
    <el-col :span="12">
      <div class="grid-content ep-bg-purple" >
        <BuySide/>
      </div>


    </el-col>
    <el-col :span="12">
      <div class="grid-content ep-bg-purple" >
        <SellSide/>
      </div>

    </el-col>
  </el-row>
<!--  -->
<!--  <el-row>-->
<!--    <el-col :span="24"><div class="grid-content ep-bg-purple-dark" ></div> 111 </el-col>-->
<!--  </el-row>-->
<!--  <el-row>-->
<!--    <el-col :span="12"><div class="grid-content ep-bg-purple" ></div> 111 </el-col>-->
<!--    <el-col :span="12"><div class="grid-content ep-bg-purple-light" ></div> 111 </el-col>-->
<!--  </el-row>-->
<!--  <el-row>-->
<!--    <el-col :span="8"><div class="grid-content ep-bg-purple" /></el-col>-->
<!--    <el-col :span="8"><div class="grid-content ep-bg-purple-light" /></el-col>-->
<!--    <el-col :span="8"><div class="grid-content ep-bg-purple" /></el-col>-->
<!--  </el-row>-->
<!--  <el-row>-->
<!--    <el-col :span="6"><div class="grid-content ep-bg-purple" /></el-col>-->
<!--    <el-col :span="6"><div class="grid-content ep-bg-purple-light" /></el-col>-->
<!--    <el-col :span="6"><div class="grid-content ep-bg-purple" /></el-col>-->
<!--    <el-col :span="6"><div class="grid-content ep-bg-purple-light" /></el-col>-->
<!--  </el-row>-->
<!--  <el-row>-->
<!--    <el-col :span="4"><div class="grid-content ep-bg-purple" /></el-col>-->
<!--    <el-col :span="4"><div class="grid-content ep-bg-purple-light" /></el-col>-->
<!--    <el-col :span="4"><div class="grid-content ep-bg-purple" /></el-col>-->
<!--    <el-col :span="4"><div class="grid-content ep-bg-purple-light" /></el-col>-->
<!--    <el-col :span="4"><div class="grid-content ep-bg-purple" /></el-col>-->
<!--    <el-col :span="4"><div class="grid-content ep-bg-purple-light" /></el-col>-->
<!--  </el-row>-->





</template>

<style scoped>

/*.result {*/
/*  height: 20px;*/
/*  line-height: 20px;*/
/*  margin: 1.5rem auto;*/
/*}*/

/*.input-box .btn {*/
/*  width: 60px;*/
/*  height: 30px;*/
/*  line-height: 30px;*/
/*  border-radius: 3px;*/
/*  border: none;*/
/*  margin: 0 0 0 20px;*/
/*  padding: 0 8px;*/
/*  cursor: pointer;*/
/*}*/

/*.input-box .btn:hover {*/
/*  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);*/
/*  color: #333333;*/
/*}*/

/*.input-box .input {*/
/*  border: none;*/
/*  border-radius: 3px;*/
/*  outline: none;*/
/*  height: 30px;*/
/*  line-height: 30px;*/
/*  padding: 0 10px;*/
/*  background-color: rgba(240, 240, 240, 1);*/
/*  -webkit-font-smoothing: antialiased;*/
/*}*/

/*.input-box .input:hover {*/
/*  border: none;*/
/*  background-color: rgba(255, 255, 255, 1);*/
/*}*/

/*.input-box .input:focus {*/
/*  border: none;*/
/*  background-color: rgba(255, 255, 255, 1);*/
/*}*/
</style>
