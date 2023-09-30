import { GetAllRatings, CreateRating } from "@wails";
import type { Rank } from "../../types/Rank";
import { safe, Safe, makeError, makeSuccess } from "../../utils";

export class RankApi {
  public async getAll(): Promise<Safe<Rank[]>> {
    const response = await safe(GetAllRatings());

    if (!response.success)
      return makeError("Error while loading your ratings :(");

    return makeSuccess(this.mapToRank(response.data));
  }

  public async create(value: number): Promise<Safe<number>> {
    const response = await safe(CreateRating(+value));

    if (!response.success) return makeError("Error while rating your day :(");

    return makeSuccess(response.data);
  }

  private mapToRank(data: Array<any> | null): Array<Rank> {
    return (
      data?.map((el) => ({
        id: el.ID,
        value: el.Value,
        createdAt: new Date(el.CreatedAt),
        updatedAt: new Date(el.UpdatedAt),
      })) || []
    );
  }

  // INFO: SINGLETON LOGIC
  private static instance: RankApi;

  private constructor() { }

  public static getInstance(): RankApi {
    if (!RankApi.instance) {
      RankApi.instance = new RankApi();
    }

    return RankApi.instance;
  }
}
