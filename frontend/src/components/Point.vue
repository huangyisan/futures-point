<template>
<!--  <el-space wrap>-->
  <el-card class="box-card">


    <el-row>
    <el-col :span="11">
<!--      <div>{{ sellSide }}</div>-->
<!--    <div v-for="(item,index) in sellItems" :key="index" class="text item">{{index}} {{ item }}</div>-->
      <el-table :data="buyTableData" style="width: 100%">
        <el-table-column prop="direction" label="方向" />
        <el-table-column prop="ratio" label="比例" />
        <el-table-column prop="price" label="价格" />
      </el-table>
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
let buyTableData = ref([
  {
    "direction": "buy",
    "ratio": "50x",
    "price": "1"
  },
  {
    "direction": "buy",
    "ratio": "50x",
    "price": "10"
  },
  {
    "direction": "buy",
    "ratio": "50x",
    "price": "100"
  },
  {
    "direction": "buy",
    "ratio": "20x",
    "price": "2"
  },
  {
    "direction": "buy",
    "ratio": "20x",
    "price": "20"
  },
  {
    "direction": "buy",
    "ratio": "20x",
    "price": "200"
  },
])

function printSide() {
  let tmpMap = []
  for (let buyItemsKey in buyItems.value) {
    console.log(buyItemsKey)

    console.log(buyItems.value[buyItemsKey].length)
    for (let i = 0; i < buyItems.value[buyItemsKey].length; i++) {

      buyTableData.value.push({
        "direction": "buy",
        "ratio": buyItemsKey,
        "price": buyItems.value[buyItemsKey][i]
      })
    }
  }
}

function makeTable(tableData,items,direction) {
  tableData.value = []
  for (let valueKey in items.value) {
    console.log(tableData.value[valueKey].length)
    for (let i = 0; i < tableData.value[valueKey].length; i++) {

      buyTableData.value.push({
        "direction": direction,
        "ratio": valueKey,
        "price": tableData.value[valueKey][i]
      })
    }
  }
}

watch(() => props.res.buyData, (newVal, oldVal) => {
  buySide.value = newVal.side
  buyItems.value = newVal.data
  // makeTable(buyTableData, buyItems, "buy")
  buyTableData.value = []
  for (let buyItemsKey in buyItems.value) {
    console.log(buyItemsKey)

    console.log(buyItems.value[buyItemsKey].length)
    for (let i = 0; i < buyItems.value[buyItemsKey].length; i++) {

      buyTableData.value.push({
        "direction": "buy",
        "ratio": buyItemsKey,
        "price": buyItems.value[buyItemsKey][i]
      })
    }
  }
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