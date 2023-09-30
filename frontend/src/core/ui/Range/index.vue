<script setup lang="ts">
import { defineProps, defineEmits } from "vue";
type Props = {
  modelValue: number | string;
  max: number;
  min: number;
};
const $emit = defineEmits(["update:modelValue"]);
const props = defineProps<Props>();
</script>
<template>
  <div class="app-range">
    <span>{{ props.modelValue }}</span>
    <input
      type="range"
      :value="props.modelValue"
      @input="
        $emit('update:modelValue', ($event.target as HTMLInputElement).value)
      "
      :max="props.max"
      :min="props.min"
    />
  </div>
</template>

<style scoped>
.app-range {
  display: grid;
  grid-template-columns: 20px 1fr;
  gap: 2rem;
  color: #c8c093;
  padding: 0.2rem 0;
}
.app-range input {
  margin: 0 1rem;
}
</style>
