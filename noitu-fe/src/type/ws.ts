export interface ResponseMessage {
  msg: string;
  currentWord: string;
  state: STATE;
  score: number;
}


export enum STATE{
  NONE = "NONE",
	QUEUEING = "QUEUEING",
	PLAYING  = "PLAYING",
  WAITING = "WAITING",
	ENDING = "ENDING",
}