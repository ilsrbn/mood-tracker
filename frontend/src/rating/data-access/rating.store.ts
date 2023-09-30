import { defineStore } from "pinia";
import { computed, reactive, ref } from "vue";
import Api from "../../core/api";
import { Rank } from "../../core/types";

// INFO: Logic shared accross multiple components
// 1. loadRatings()
// 2. canAddNewRank
export const useRatingStore = defineStore("rating", () => {
  // INFO: initialazation
  const _rankApi = Api.Ranking;

  const ratings = ref<Rank[]>([]);

  const loading = ref(true);
  const error = reactive({
    state: false,
    message: "",
  });

  // INFO: Base utulities and computed
  const _setRatings = (value: Rank[]) => {
    ratings.value = value;
  };

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

  const latestRanking = computed(() => {
    if (ratings.value.length < 1) return null;
    return ratings.value[ratings.value.length - 1];
  });

  const _isFirstDateLarger = (date1: Date, date2: Date) => {
    const time1 = date1.setHours(0, 0, 0, 0);
    const time2 = date2.setHours(0, 0, 0, 0);

    return time1 > time2;
  };

  // INFO: Business logic
  const loadRatings = async () => {
    _setLoading(true);
    const response = await _rankApi.getAll();
    if (!response.success) {
      _setError(response.error);
    } else {
      _resetError();
      _setRatings(response.data);
    }
    _setLoading(false);
  };

  const canAddNewRank = computed(() => {
    if (loading.value) return false;
    if (!latestRanking.value) return true;

    const today = new Date();
    const latestRankedDate = latestRanking.value.createdAt;

    return _isFirstDateLarger(today, latestRankedDate);
  });

  return { ratings, loading, error, loadRatings, canAddNewRank, latestRanking };
});
