<script setup lang="ts">
import LineChart from "../../core/ui/LineChart/index.vue";
import { computed, defineProps } from "vue";
import type { Rank } from "../../core/types/Rank";

type Props = {
  ratings: Rank[];
};

const props = defineProps<Props>();

const localized = (date: Date) => {
  return date.toLocaleDateString("en-EN", {
    year: "numeric",
    month: "short",
    day: "numeric",
  });
};

const formattedRatings = computed(() => {
  return props.ratings.map((el) => ({
    ...el,
    createdAt: localized(el.createdAt),
    updatedAt: localized(el.updatedAt),
  }));
});
</script>
<template>
  <section>
    <LineChart
      :data="formattedRatings"
      x="createdAt"
      y="value"
      label="Day's rating"
    />
  </section>
</template>
<style scoped>
section {
  margin: 2rem auto 0;
  max-width: 100%;
}
</style>
