<script setup lang="ts">
import { computed, onMounted } from "vue";
import Form from "../ui/rating-form.vue";
import Chart from "../ui/rating-chart.vue";
import AlreadyRated from "../ui/ranking-already-rated.vue";
import { useRatingStore } from "../data-access/rating.store";
import { useRatingService } from "../data-access/rating.service";

import Alert from "../../core/ui/Alert/index.vue";

const store = useRatingStore();
const service = useRatingService();

const loading = computed(() => store.loading);
const error = computed(() => store.error);
const ratings = computed(() => store.ratings);
const latestRatedValue = computed(() => store.latestRanking?.value || 0);

onMounted(() => {
  store.loadRatings();
});

const canAddNewRank = computed(() => store.canAddNewRank);

const onSubmit = (value: number) => {
  service.rate(value);
};

const showChart = computed(() => error.value.state !== true);
</script>
<template>
  <article id="rating-shell">
    <main>
      <Transition mode="out-in">
        <Form v-if="canAddNewRank" @submit="onSubmit" />
        <AlreadyRated v-else :latest-ranking-value="latestRatedValue" />
      </Transition>
    </main>
    <Transition>
      <footer v-if="!loading">
        <Chart v-if="showChart" :ratings="ratings" />
        <Alert v-else v-model="error.state" type="error">{{
          error.message
        }}</Alert>
      </footer>
    </Transition>
  </article>
</template>
<style scoped>
article#rating-shell main,
article#rating-shell footer {
  margin: 2rem 0;
}

.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>
