import { randomUUID } from "node:crypto";

class Salon {
  id: string;
  name: string;
  openTime: Date;
  closeTime: Date;
  createdAt: Date;

  constructor(
    name: string,
    openTime: string = "09:00",
    closeTime: string = "18:00"
  ) {
    this.id = randomUUID();
    this.name = name;
    this.openTime = this.parseTimeString(openTime);
    this.closeTime = this.parseTimeString(closeTime);
    this.createdAt = new Date();
  }

  // Helper function to parse time strings into Date objects
  private parseTimeString(timeString: string): Date {
    const [hours, minutes] = timeString.split(":").map(Number);
    const date = new Date();
    date.setHours(hours, minutes, 0, 0);
    return date;
  }
}

const s = new Salon("Fio");
console.log(s);
