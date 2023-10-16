<template>
<!--  <el-space wrap>-->
  <el-card class="box-card">
    <el-row>
    <el-col :span="11">
      <el-table
          :data="buyTableData"
          height="250"
          style="width: 100%"
          :row-class-name="tableRowClassName"
          stripe
      >
        <el-table-column prop="ratio" label="比例" fixed/>
        <el-table-column prop="price" label="价格" />
        <el-table-column prop="direction" label="方向" />
      </el-table>
    </el-col>
      <el-col :span="2">
        <el-divider direction="vertical" content-position="center" style="height: 100%"/>
      </el-col>
    <el-col :span="11">
      <el-table
          :data="sellTableData"
          height="250"
          style="width: 100%"
          max-height="250"
          :row-class-name="tableRowClassName"
          stripe
      >

        <el-table-column prop="ratio" label="比例" fixed/>
        <el-table-column prop="price" label="价格" />
        <el-table-column prop="direction" label="方向" />
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
let buyTableData = ref()
let sellTableData = ref([])

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

// const ratioFilter = () => {
//   for (let buyItemsKey in buyItems.value) {
//     // console.log(buyItemsKey)
//     if (ratios.value.indexOf(buyItemsKey) === -1) {
//       ratios.value.push(buyItemsKey)
//     }
//   }
//   console.log(ratios)
// }

function tableRowClassName() {
  return "warning-row"
}
// const tableRowClassName = ({row, rowIndex}) => {
//   if (row.ratio === "ratio50") {
//     return "warning-row"
//   } else {
//     return "success-row"
//   }
// }


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

.el-table .warning-row {
  --el-table-tr-bg-color: var(--el-color-warning-light-9);
}
.el-table .success-row {
  --el-table-tr-bg-color: var(--el-color-success-light-9);
}
</style>