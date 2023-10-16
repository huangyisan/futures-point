<template>
<!--  <el-space wrap>-->
  <el-card class="box-card">


    <el-row>
    <el-col :span="11">
      <div>{{ sellSide }}</div>
    <div v-for="(item,index) in sellItems" :key="index" class="text item">{{index}} {{ item }}</div>
    </el-col>
      <el-col :span="2">

        <el-divider direction="vertical" content-position="center" />
      </el-col>
    <el-col :span="11">
      <div>{{ buySide }}</div>
    <div v-for="(item,index) in buyItems" :key="index" class="text item">{{index}} {{ item }}</div>
    </el-col>
    </el-row>
    <el-divider direction="horizontal" />
    <button @click="printSide">1111</button>
  </el-card>
<!--  </el-space>-->
  <!-- -->
</template>

<script setup>
import {watch, ref} from "vue";

const props = defineProps(['res'])
let buyItems = ref([])
let buySide = ref()
let sellItems = ref([])
let sellSide = ref()
// let buyItems = props.res.buyData;
// let sellItems = props.res.sellData;
// let buySide = ref(props.res.buyData.side)
// let sellSide = ref(props.res.sellData.side)
// let tmp = ref()
function printSide() {
  // console.log(buyItems.value)
  // console.log(props.res.sellData)

}

watch(() => props.res.buyData, (newVal, oldVal) => {
  buySide.value = newVal.side
  buyItems.value = newVal.data
})

watch(() => props.res.sellData, (newVal, oldVal) => {
  console.log(`new: ${newVal.side}`)
  sellSide.value = newVal.side
  sellItems.value = newVal.data
})



</script>

<style scoped>
.text {
  font-size: 14px;
}

.item {
  padding: 18px 0;
}

/*.box-card {*/
/*  width: 480px;*/
/*}*/
</style>