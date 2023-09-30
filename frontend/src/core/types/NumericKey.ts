export type NumericKey<T> = {
  [K in keyof T]: T[K] extends number | bigint ? K : never;
}[keyof T];
