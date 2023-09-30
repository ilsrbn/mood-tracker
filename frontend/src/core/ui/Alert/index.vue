<script setup lang="ts">
import { defineProps, defineEmits, computed } from "vue";
type Props = {
  modelValue: boolean;
  type: "success" | "error";
};

const props = defineProps<Props>();
const $emit = defineEmits(["update:modelValue"]);

const icon = computed(() => {
  switch (props.type) {
    case "success":
      return "done";
    case "error":
    default:
      return "error";
  }
});

const onClose = () => {
  $emit("update:modelValue", false);
};
</script>

<template>
  <section v-if="props.modelValue" :class="[props.type]">
    <header>
      <span class="material-symbols-outlined" @click="onClose"> close </span>
    </header>
    <main>
      <span class="material-symbols-outlined"> {{ icon }} </span>
      <div>
        <slot></slot>
      </div>
    </main>
  </section>
</template>
<style scoped>
* {
  --color-success: #215231;
  --color-error: #522421;
}

section {
  margin: 1rem auto;
  width: 90%;
  max-width: 90%;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  border-radius: 20px;

  position: relative;
}

header {
  position: absolute;
  right: 0px;
  top: 0px;
  padding: 0.5rem;
}

header span {
  cursor: pointer;
}

main {
  padding: 0.5rem;
  display: grid;
  grid-template-columns: 20px 1fr;
  gap: 20px;
}

.success {
  background-color: var(--color-success);
}

.error {
  background-color: var(--color-error);
}
</style>
