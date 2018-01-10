pecheck\_simple is a simple Job Submission Verifier (JSV) for Univa Grid Engine.
It is written in [Go](https://golang.org) using [Daniel Gruber's jsv package](https://github.com/dgruber/jsv). See also: http://www.gridengine.eu/index.php/grid-engine-internals/171-main-memory-limitation-with-grid-engine-a-short-introduction-into-cgroups-20130825

Things to note:
0. This JSV tries to prevent multithreaded jobs which were submitted as serial jobs from consuming CPU cores. It does this using cgroups cpuset feature. This has been tested in UGE 8.4.4 -- see the UGE Administrator's Guide sec. 2.5.6
1. This only looks for the *lack* of a PE request; if someone requests "-pe something 1", that would go through without having a binding set.

