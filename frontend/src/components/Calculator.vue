<script setup>
import {reactive, ref} from 'vue'
import {Calculate, Greet} from '../../wailsjs/go/main/App'
import Point from "./Point.vue";

const point = reactive({
  price: 0,
  buyData: {},
  sellData: {},
})

let display = ref(false)

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
    point.sellData = jsonRes
    console.log(jsonRes)

    const decodeData2 = atob(result[1])
    const jsonRes2 = JSON.parse(decodeData2)
    point.buyData = jsonRes2
    console.log(jsonRes2)
  })
  display.value = true
}

</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-main>
        <el-row class="center" :gutter="10">
          <el-col :span="24">
            <div v-if="display">
              <Point :res="point"/>
            </div>
            <div v-else>
              <el-empty description="No data"/>
            </div>
          </el-col>
        </el-row>
        <el-row class="center" :gutter="5">
          <el-col class="grid-content" :span="12">
            <el-text size="small">
              BUY
            </el-text>
          </el-col>
          <el-col class="grid-content" :span="12">
            <el-text size="small">
              SELL
            </el-text>
          </el-col>
        </el-row>
        <el-row class="center" :gutter="5">
          <el-col :span="1" :offset=16>
            <el-input-number controls-position="right" class="item" v-model="point.price" @change="calculate" :min="0"
                             type="number" size="small"/>
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
.grid-content {
border-radius: 4px;
  min-height: 36px;
  padding: 5px;
  color: #fff;
  text-align: center;

}

</style>
