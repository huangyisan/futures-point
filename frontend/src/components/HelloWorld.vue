<script setup>
import {reactive} from 'vue'
import {Calculate, Greet} from '../../wailsjs/go/main/App'
import Point from "./Point.vue";

const point = reactive({
  price: 0,
  buyData: {
  },
  sellData: {
  },
})

// const pconst point = reactive({
//   price: 0,
//   buyRes: "",
//   sellRes: ""
// })oint = reactive({
//   price: 0,
//   buyRes: "",
//   sellRes: ""
// })

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}

function calculate() {
  Calculate(point.price).then(result => {
    const decodeData = atob(result[0])
    const jsonRes = JSON.parse(decodeData)
    point.sellData =  jsonRes
    console.log(jsonRes)

    const decodeData2 = atob(result[1])
    const jsonRes2 = JSON.parse(decodeData2)
    point.buyData =  jsonRes2
    console.log(jsonRes2)
  })
}

</script>

<template>
  <div class="common-layout">
    <el-container>

      <el-header>
        <el-row class="center">
          <el-col :span="24">
            <!--        <div id="result">当前价格</div>-->
            <!--        <div id="input" class="input-box">-->
            <el-input-number v-model="point.price" :min="0" type="number"/>
            <!--          <input id="name" v-model="point.price" autocomplete="off" class="input" type="number"/>-->
            <el-button type="primary" @click="calculate">计算</el-button>
            <!--          <button class="btn" @click="calculate">计算</button>-->
            <!--        </div>-->
          </el-col>
        </el-row>
      </el-header>
      <el-main>
        <el-row class="center">
          <el-col :span="24">
            <div>
              <Point :res="point"/>
<!--              <el-divider direction="vertical" />-->
            </div>
          </el-col>
        </el-row>
      </el-main>

    </el-container>
  </div>

</template>

<style scoped>
.center {
  text-align: center;
}
</style>
