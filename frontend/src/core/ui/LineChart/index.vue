<script setup lang="ts">
import Chart from "chart.js/auto";
import H1 from "../Header/h1.vue";
import { onMounted, ref, watch } from "vue";
import { defineProps } from "vue";

interface Props {
  data: Array<any>;
  x: string;
  y: string;
  label: string;
}
const props = defineProps<Props>();

const makeChart = () => {
  new Chart(document.getElementById("chart") as HTMLCanvasElement, {
    type: "line",
    data: {
      labels: props.data.map((row) => row[props.x]),
      datasets: [
        {
          label: props.label,
          data: props.data.map((row) => row[props.y]),
          borderColor: "#c8c093",
          fill: false,
          cubicInterpolationMode: "monotone",
          tension: 0.4,
        },
      ],
    },
    options: {
      scales: {
        y: {
          min: 1,
          max: 10,
        },
      },
    },
  });
};

onMounted(() => makeChart());

watch(
  () => props,
  () => makeChart(),
  { deep: true },
);
</script>
<template>
  <H1>Days rated: {{ props.data.length }}</H1>
  <canvas style="margin-top: 2rem" id="chart" />
</template>
