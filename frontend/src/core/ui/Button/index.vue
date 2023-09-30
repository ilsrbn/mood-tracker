<script setup lang="ts"></script>
<template>
  <button v-bind="$attrs">
    <slot />
  </button>
</template>
<style scoped>
button {
  margin: 0;
  padding: 0;
  border: none;
  outline: none;
  box-shadow: none;

  background: #c8c093;
  color: #1f1f28;
  padding: 0.5rem 2rem;
  border-radius: 5px;
  font-weight: bold;
  transition: all 150ms linear;
  cursor: pointer;
}

button:hover {
  background: #cccccc;
}

button:focus-visible {
  outline: 2px solid #c8c093 !important;
  border: 2px solid #1f1f28;
  /* border-radius: 3px; */
}
</style>
