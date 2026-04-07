import {
  startOfMonth,
  endOfMonth,
  startOfWeek,
  endOfWeek,
  eachDayOfInterval,
  format,
  isSameMonth,
  isToday,
  isSameDay,
  isWithinInterval,
} from "date-fns";

export function getCalendarDays(year: number, month: number): Date[] {
  const monthStart = startOfMonth(new Date(year, month));
  const monthEnd = endOfMonth(monthStart);
  const calStart = startOfWeek(monthStart, { weekStartsOn: 0 }); // Sunday
  const calEnd = endOfWeek(monthEnd, { weekStartsOn: 0 });

  return eachDayOfInterval({ start: calStart, end: calEnd });
}

export function formatDateKey(date: Date): string {
  return format(date, "yyyy-MM-dd");
}

export function isInRange(
  date: Date,
  start: Date | null,
  end: Date | null
): boolean {
  if (!start || !end) return false;
  const from = start <= end ? start : end;
  const to = start <= end ? end : start;
  return isWithinInterval(date, { start: from, end: to });
}

export { isSameMonth, isToday, isSameDay, format };
