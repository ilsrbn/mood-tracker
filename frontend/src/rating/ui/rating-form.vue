<script lang="ts" setup>
import { ref, computed, defineEmits } from "vue";
import { useRatingService } from "../data-access/rating.service";

import Range from "../../core/ui/Range/index.vue";
import Button from "../../core/ui/Button/index.vue";
import H1 from "../../core/ui/Header/h1.vue";
import Alert from "../../core/ui/Alert/index.vue";

type Emits = {
  (e: "submit", value: number): void;
};

const $emits = defineEmits<Emits>();

const rating = ref(10);

const ratingService = useRatingService();

const loading = computed(() => ratingService.loading);
const error = computed(() => ratingService.error);
const config = ratingService.config;

const onSubmit = () => $emits("submit", +rating.value);
</script>

<template>
  <form
    fast-fail
    @submit.prevent="onSubmit"
    :loading="loading"
    :disabled="loading"
  >
    <Alert v-model="error.state" type="error">{{ error.message }}</Alert>
    <H1>Rate your day :)</H1>
    <Range v-model="rating" :max="config.max" :min="config.min" />
    <Button type="submit" style="margin-top: 1rem">Rate</Button>
  </form>
</template>

<style scoped></style>
