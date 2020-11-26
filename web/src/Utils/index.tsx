export const humanizedDatetime = (second: number) => {
  return new Date(second * 1000).toLocaleString("zh-CN");
};
