export interface CalendarNote {
  id: string;
  dateKey: string; // "YYYY-MM-DD"
  text: string;
  createdAt: string;
}

export interface DateRange {
  start: Date | null;
  end: Date | null;
}

export interface MonthConfig {
  month: number; // 0-indexed
  year: number;
  heroImage: string;
  heroAlt: string;
}
