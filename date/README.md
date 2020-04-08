# format/date

Date formatting based on unicode specification
https://www.unicode.org/reports/tr35/tr35-dates.html#Date_Field_Symbol_Table
Inspired by https://github.com/date-fns/date-fns

```go
import (
  "fmt"
  "github.com/datasweet/format/date"
)

func main() {
  // The numeric time represents Thu Feb  4 21:00:57.012345600 PST 2009
  tm := time.Unix(0, 1233810057012345600)
  
  // Format with local timezone Europe/Paris
  // Pattern "yyyy-MM-ddTHH:mm:ssXXX" = RFC3339
  fmt.Println(date.Format(language.French, tm, "yyyy-MM-ddTHH:mm:ssXXX"))
  // Output: 2009-02-05T06:00:57+01:00
}
```

## Patterns

*	_a_: AM, PM
* _b_: AM, PM, noon, midnight
* _c_: Stand-alone local day of week <!> not implemented
* _d_: Day of month
* _e_: Local day of week
* _g_: Modified Julian day <!> not implemented
* _h_: Hour [1-12]
* _j_: Localized hour w/ day period <!> not implemented
* _k_: Hour [1-24]
* _m_: Minute		
* _q_: Stand-alone quarter
* _r_: Related Gregorian year  <!> not implemented
* _s_: Second
* _t_: timestamp in millisecond custom  <!> custom code
* _u_: Extended year
* _v_: Timezone (generic non-locat.) <!> not implemented
* _w_: Local week of year  <!> not implemented
* _x_: Timezone (ISO-8601 w/o Z)
* _y_: Year (abs)
* _z_: Timezone (specific non-locat.)
* _A_: Milliseconds in day <!> not implemented
* _B_: Flexible day period		
* _C_: Localized hour w/ day period <!> not implemented
* _D_: Day of year
* _E_: Day of week
* _F_: Day of week in month  <!> not implemented
* _G_: Era
* _H_: Hour [0-23]
* _J_: Localized hour w/o day period <!> not implemented
* _K_: Hour [0-11]
* _L_: Stand-alone month
* _M_: Month
* _O_: Timezone (GMT)
* _Q_: Quarter
* _S_: Fraction of second
* _U_: Cyclic year  <!> not implemented
* _V_: Timezone (location) <!> not implemented
* _W_: Week of month  <!> not implemented
* _X_: Timezone (ISO-8601)
* _Y_: Local week-numbering year
* _Z_: Timezone (aliases)
