'reach 0.1';

export const main =
  Reach.App(
    {
      deployMode: 'firstMsg',
      connectors: [ETH],
    },
    [['Alice', { request: UInt,
                 info: Bytes }],
     [  'Bob', { want: Fun([UInt], Null),
                 got: Fun([Bytes], Null) }]],
    (A, B) => {
      A.only(() => {
        const request = declassify(interact.request); });
      A.publish(request);
      commit();

      B.only(() => {
        interact.want(request); });
      B.pay(request);
      commit();

      A.only(() => {
        const info = declassify(interact.info); });
      A.publish(info);
      transfer(request).to(A);
      commit();

      B.only(() => {
        interact.got(info); });
      exit();
    } );
