<template>
<!--  <el-space wrap>-->
  <el-card class="box-card">


    <el-row>
    <el-col :span="11">
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
      <el-table :data="sellTableData" style="width: 100%">
        <el-table-column prop="direction" label="方向" />
        <el-table-column prop="ratio" label="比例" />
        <el-table-column prop="price" label="价格" />
      </el-table>
    </el-col>
    </el-row>
    <el-divider direction="horizontal" />
    <button @click="printSide">1111</button>
  </el-card>
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
let sellTableData = ref([
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

const makeTable = (tableData, direction, items) => {
  tableData.value = []
  for (let itemsKey in items.value) {
    for (let i = 0; i < items.value[itemsKey].length; i++) {
      tableData.value.push({
        "direction": direction,
        "ratio": itemsKey,
        "price": items.value[itemsKey][i]
      })
    }
  }
}

watch(() => props.res.buyData, (newVal, oldVal) => {
  buySide.value = newVal.side
  buyItems.value = newVal.data
  makeTable(buyTableData, "buy", buyItems)
})

watch(() => props.res.sellData, (newVal, oldVal) => {
  sellSide.value = newVal.side
  sellItems.value = newVal.data
  makeTable(sellTableData, "sell", sellItems)
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