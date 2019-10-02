<h1> CetCest </h1>
<p> CetCest is a small library to find out if a timestamp is CET or CEST. <br />
This is ONLY handy if you receive timestamps which does not include info regarding
CET/CEST, but you know the time is either CET or CEST.  <br /> <br />
You can either chose to find out if a timestamp is CET (IsCet()), or you can 
receive the entire timestamp with either CET or CEST appended</p>


<h2> How to use </h2>
<p> Import:  <br />
cetcest "github.com/TobiasFP/TimeInCetOrCestToEpoch_GoLang/CetCest"
<br />
Use: <br />
give input of the format: <br />
layout := "2006-01-02T15:04:00"
cetcest.IsCet(layout) // Returns true
<br />

summertime := "2018-04-01T07:45:00"
summertime = cetcest.GetCETOrCEST(summertime) // returns "2018-04-01T07:45:00 CEST"
<br />
It is now simple to get the epoch timestamp as follows:
, err := time.Parse(layout, summertime)
if err != nil {
    fmt.Println(err)
    t.Error("Didnt Expect to receive: Expected: a time'")

}
</p>

<h2> Problems: </h2>
<p> Since this was a really quick fix to my problem, and I don't care about
the two hours overlap at night during last sunday of march and oktober, there
are some problems in regard to this. <br /> 
If you need a fix for this,  simply fix the switch/case in IsCET().

 </p>