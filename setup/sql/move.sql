CREATE table temp (like spreads);

INSERT into temp
  (SELECT
    "time" as ts,
     dvbid as bid,
     dvask as ask,
     size,
     dvbps as width_bps,
     ticker,
     'DV'
    from quote_widths
  );


INSERT into temp
  (SELECT
    "time" as ts,
     enbid as bid,
     enask as ask,
     size,
     enbps as width_bps,
     ticker,
     'Enigma'
    from quote_widths
    where enbid != -1.0
  );

-- INSERT into temp (SELECT * from spreads);

INSERT into oldspreads (SELECT * from temp order by ts,lp,ticker,size);

DROP table temp;
