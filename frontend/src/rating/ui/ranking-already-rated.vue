<script setup lang="ts">
import { defineProps } from "vue";
import H1 from "../../core/ui/Header/h1.vue";

type Props = {
  latestRankingValue: number;
};

const props = defineProps<Props>();
</script>
<template>
  <H1
    >You have already rated your day as "{{ props.latestRankingValue }}/10"!</H1
  >
</template>
<style scoped></style>
