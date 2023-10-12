<script setup>
import {reactive} from 'vue'
import {Greet, Calculate} from '../../wailsjs/go/main/App'
import BuySide from "./BuySide.vue";
import SellSide from "./SellSide.vue";

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})

const point = reactive({
  price: 0,
  buyRes:"",
  sellRes: ""
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}

function calculate() {
  Calculate(point.price).then(result => {
    point.buyRes = result[0]
    point.sellRes = result[1]
  })
}

</script>

<template>

    <el-row class="center">
      <el-col :span="24">
<!--        <div id="result">当前价格</div>-->
<!--        <div id="input" class="input-box">-->
          <el-input-number v-model="point.price" :min="0" type="number" />
<!--          <input id="name" v-model="point.price" autocomplete="off" class="input" type="number"/>-->
        <el-button type="primary" @click="calculate">计算</el-button>
<!--          <button class="btn" @click="calculate">计算</button>-->
<!--        </div>-->
      </el-col>
    </el-row>
  <el-row class="center">
    <el-col :span="12">
      <div >
        <BuySide :buyRes="point.buyRes"/>
      </div>
    </el-col>
    <el-col :span="12">
      <div >
        <SellSide :sellRes="point.sellRes"/>
      </div>
    </el-col>
  </el-row>


</template>

<style scoped>
.center {
  text-align: center;
}
</style>
