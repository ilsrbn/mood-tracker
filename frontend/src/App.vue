<script lang="ts" setup>
import RatingShell from "./rating/feature/rating-shell.vue";
</script>

<template>
  <article>
    <span style="height: 20%"></span>
    <RatingShell />
  </article>
</template>

<style scoped>
article {
  height: 100svh;
  width: 100%;
  max-width: 960px;
  margin: 0 auto;
  padding: 0 2rem;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}
</style>
