import { reactive, ref } from "vue";
import Api from "../../core/api";
import { useRatingStore } from "./rating.store";

// INFO: Local module logic only
// 1. rate()
export const useRatingService = () => {
  // INFO: initialazation
  const _rankApi = Api.Ranking;
  const _store = useRatingStore();

  const config = Object.freeze({
    max: 10,
    min: 1,
  });

  const loading = ref(true);
  const error = reactive({
    state: false,
    message: "",
  });

  // INFO: Base utulities and computed
  const _setLoading = (value: boolean) => {
    loading.value = value;
  };

  const _setError = (message: string) => {
    error.message = message;
    error.state = true;
  };

  const _resetError = () => {
    error.message = "";
    error.state = false;
  };

  // INFO: Business logic
  const rate = async (value: number) => {
    _setLoading(true);
    const response = await _rankApi.create(value);
    if (!response.success) {
      _setError(response.error);
    } else {
      _resetError();
      _store.loadRatings();
    }
    _setLoading(false);
  };

  return { rate, config, loading, error };
};
