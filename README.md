literate-computing-machine
==========================

This is a quick and dirty project to illustrate a bug where Golang `time.After` fails to wake the process for large delays.

    3h0m0s 2016-10-20 10:59:04.233645865 +0100 BST: should wake at 2016-10-20 13:59:04.233645865 +0100 BST in 3h0m0s
    9h0m0s 2016-10-20 10:59:04.350703594 +0100 BST: should wake at 2016-10-20 19:59:04.350703594 +0100 BST in 9h0m0s
    10m0s 2016-10-20 10:59:04.350711442 +0100 BST: should wake at 2016-10-20 11:09:04.350711442 +0100 BST in 10m0s
    8h0m0s 2016-10-20 10:59:04.365249689 +0100 BST: should wake at 2016-10-20 18:59:04.365249689 +0100 BST in 8h0m0s
    1h0m0s 2016-10-20 10:59:04.385879887 +0100 BST: should wake at 2016-10-20 11:59:04.385879887 +0100 BST in 1h0m0s
    7h0m0s 2016-10-20 10:59:04.389250374 +0100 BST: should wake at 2016-10-20 17:59:04.389250374 +0100 BST in 7h0m0s
    2h0m0s 2016-10-20 10:59:04.402031663 +0100 BST: should wake at 2016-10-20 12:59:04.402031663 +0100 BST in 2h0m0s
    5h0m0s 2016-10-20 10:59:04.411631508 +0100 BST: should wake at 2016-10-20 15:59:04.411631508 +0100 BST in 5h0m0s
    4h0m0s 2016-10-20 10:59:04.44162023 +0100 BST: should wake at 2016-10-20 14:59:04.44162023 +0100 BST in 4h0m0s
    10s 2016-10-20 10:59:04.442757967 +0100 BST: should wake at 2016-10-20 10:59:14.442757967 +0100 BST in 10s
    6h0m0s 2016-10-20 10:59:04.446185807 +0100 BST: should wake at 2016-10-20 16:59:04.446185807 +0100 BST in 6h0m0s
    10s 2016-10-20 10:59:14.442956722 +0100 BST: did wake with delta 198.755µs

    real	0m10.850s
    user	0m0.169s
    sys	0m0.099s
    10m0s 2016-10-20 11:09:04.351541782 +0100 BST: did wake with delta 830.34µs

    real	10m0.759s
    user	0m0.165s
    sys	0m0.090s
    1h0m0s 2016-10-20 11:59:04.389439653 +0100 BST: did wake with delta 3.559766ms

    real	60m0.799s
    user	0m0.198s
    sys	0m0.093s
    2h0m0s 2016-10-20 12:59:04.406737243 +0100 BST: did wake with delta 4.70558ms

    real	120m0.816s
    user	0m0.240s
    sys	0m0.111s
    3h0m0s 2016-10-20 13:59:04.235482729 +0100 BST: did wake with delta 1.836864ms

    real	180m0.644s
    user	0m0.300s
    sys	0m0.130s
    4h0m0s 2016-10-20 14:59:04.443901762 +0100 BST: did wake with delta 2.281532ms

    real	240m0.854s
    user	0m0.322s
    sys	0m0.119s
    5h0m0s 2016-10-20 15:59:04.415292046 +0100 BST: did wake with delta 3.660538ms

    real	300m0.823s
    user	0m0.360s
    sys	0m0.127s
    6h0m0s 2016-10-20 17:26:37.232606209 +0100 BST: did wake with delta 27m32.786420402s

    real	387m33.638s
    user	0m0.436s
    sys	0m0.149s
    7h0m0s 2016-10-20 18:26:37.106103255 +0100 BST: did wake with delta 27m32.716852881s

    real	447m33.502s
    user	0m0.439s
    sys	0m0.144s
    9h0m0s 2016-10-21 09:37:21.668181946 +0100 BST: did wake with delta 13h38m17.317478352s
    8h0m0s 2016-10-21 09:37:21.668608365 +0100 BST: did wake with delta 14h38m17.303358676s

    real	1358m18.063s
    user	0m0.459s
    sys	0m0.142s

    real	1358m18.049s
    user	0m0.469s
    sys	0m0.148s


Possibly caused by app nap? Resizing the terminal window (which I believe results in signals being sent) unfroze the 8h and 9h delays which failed to fire overnight (see above).