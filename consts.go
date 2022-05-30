package main

const RESET string = "\033[0m"
const BOLD_GREEN string = "\033[1;32m"
const GREEN string = "\033[32m"
const RED string = "\033[31m"

const XINITRC string = `c2V0eGtibWFwIC1tb2RlbCBhYm50MiAtbGF5b
3V0IGJyCnN4aGtkICYKcGljb20gJgpuaXRyb2dlbiAtLXJlc3RvcmUgJgp4c2
V0cm9vdCAtY3Vyc29yX25hbWUgbGVmdF9wdHIKZXhlYyBic3B3bQ==`

const BSPWMRC string = `IyEgL2Jpbi9zaAoKJEhPTUUvLmNvbmZpZy9wb
2x5YmFyL2xhdW5jaC5zaAoKYnNwYyBtb25pdG9yIC1kIEkgSUkgSUlJIElWIF
YgVkkgVklJIFZJSUkgSVggWAoKYnNwYyBjb25maWcgYm9yZGVyX3dpZHRoICA
gICAgICAgMQpic3BjIGNvbmZpZyB3aW5kb3dfZ2FwICAgICAgICAgIDEyCgpi
c3BjIGNvbmZpZyBzcGxpdF9yYXRpbyAgICAgICAgICAwLjUKYnNwYyBjb25ma
Wcgbm9ybWFsX2JvcmRlcl9jb2xvciAiI2ZmMDAwMCIKYnNwYyBjb25maWcgZm
9jdXNlZF9ib3JkZXJfY29sb3IgIiMwMGZmMDAiCmJzcGMgY29uZmlnIGZvY3V
zX2ZvbGxvd3NfcG9pbnRlciB0cnVlCgpic3BjIGNvbmZpZyBib3JkZXJsZXNz
X21vbm9jbGUgICB0cnVlCmJzcGMgY29uZmlnIGdhcGxlc3NfbW9ub2NsZSAgI
CAgIHRydWUKCmJzcGMgcnVsZSAtYSBHaW1wIGRlc2t0b3A9J144JyBzdGF0ZT
1mbG9hdGluZyBmb2xsb3c9b24KYnNwYyBydWxlIC1hIENocm9taXVtIGRlc2t
0b3A9J14yJwpic3BjIHJ1bGUgLWEgbXBsYXllcjIgc3RhdGU9ZmxvYXRpbmcK
YnNwYyBydWxlIC1hIEt1cGZlci5weSBmb2N1cz1vbgpic3BjIHJ1bGUgLWEgU
2NyZWVua2V5IG1hbmFnZT1vZmYK`

const RESIZE_BSPWM string = `IyEvdXNyL2Jpbi9lbnYgYmFzaAojIHJl
c2l6ZSB7dXAsZG93bixlYXN0LHdlc3R9IGRlbHRhCgpkZWx0YT0kezI6LSIzM
CJ9CgpjYXNlICQxIGluCiAgcmlnaHQpICBkaW09dzsgc2lnbj0rIDs7CiAgbG
VmdCkgICBkaW09dzsgc2lnbj0tIDs7CiAgdXApICAgICBkaW09aDsgc2lnbj0
tIDs7CiAgZG93bikgICBkaW09aDsgc2lnbj0rIDs7CiAgKikgZWNobyAiVXNh
Z2U6IHJlc2l6ZSB7dXAsZG93bixsZWZ0LHJpZ2h0fSBbZGVsdGFdIiAmJiBle
Gl0IDEgOzsKZXNhYwoKeD0wOyB5PTA7CmNhc2UgJGRpbSBpbgogIHcpIHg9JH
NpZ24kZGVsdGE7ICBkaXI9cmlnaHQ7ICBmYWxsZGlyPWxlZnQgICA7OwogIGg
pIHk9JHNpZ24kZGVsdGE7ICBkaXI9dG9wOyAgICBmYWxsZGlyPWJvdHRvbSA7
Owplc2FjCgpic3BjIG5vZGUgLXogIiRkaXIiICIkeCIgIiR5IiB8fCBic3BjI
G5vZGUgLXogIiRmYWxsZGlyIiAiJHgiICIkeSI7Cg==`

const SXHKDRC string = `IwojIHdtIGluZGVwZW5kZW50IGhvdGtleXMKI
woKIyBUZXJtaW5hbCBlbXVsYXRvcgphbHQgKyBSZXR1cm4KCXhmY2U0LXRlcm
1pbmFsCgojIFByb2dyYW0gbGF1bmNoZXIKYWx0ICsgcwoJcm9maSAtc2hvdyB
ydW4KCiMgTWFrZSBzeGhrZCByZWxvYWQgaXRzIGNvbmZpZ3VyYXRpb24gZmls
ZXM6CmFsdCArIEVzY2FwZQoJcGtpbGwgLVVTUjEgLXggc3hoa2QKCiMgTG9ja
3NjcmVlbgphbHQgKyBsCglpM2xvY2sgLWMgIiMyMDIwMjAiCgojIFBvd2Vyb2
ZmCmFsdCArIHAKCXBvd2Vyb2ZmCgojCiMgQlNQV00KIwoKIyBDbG9zZSBvciB
raWxsIG5vZGUKYWx0ICsge18sIHNoaWZ0fSArIHEKCWJzcGMgbm9kZSAte2Ms
IGt9CgojIENsb3NlIGFsbCBub2RlcyBpbiBjdXJyZW50IHdvcmtzcGFjZQphb
HQgKyBEZWxldGUKCWJzcGMgbm9kZSAnYW55LmxvY2FsJyAtLWNsb3NlCgojIF
F1aXQgb3IgcmVzdGFydCBic3B3bQphbHQgKyBjdHJsICsge3EsIHJ9Cglic3B
jIHtxdWl0LCB3bSAtcn0KCiMgQWx0ZXJuYXRlIGJldHdlZW4gdGhlIHRpbGVk
IGFuZCBtb25vY2xlIGxheW91dAphbHQgKyBUYWIKCWJzcGMgZGVza3RvcCAtb
CBuZXh0CgojIEZvY3VzIHRoZSBub2RlIGluIHRoZSBnaXZlbiBkaXJlY3Rpb2
4KYWx0ICsge0xlZnQsIERvd24sIFVwLCBSaWdodH0KCWJzcGMgbm9kZSAtZiB
7d2VzdCwgc291dGgsIG5vcnRoLCBlYXN0fQoKIyBSZXNpemUgbm9kZQphbHQg
KyBjdHJsICsge0xlZnQsIERvd24sIFVwLCBSaWdodH0KCX4vLmNvbmZpZy9ic
3B3bS9zY3JpcHRzL3Jlc2l6ZS5zaCB7bGVmdCwgZG93biwgdXAsIHJpZ2h0fS
AzMAoKIyBGb2N1cyBlc3BlY2lmaWMgZGVza3RvcAphbHQgKyB7MS05LDB9Cgl
ic3BjIGRlc2t0b3AgLWYgJ157MS05LDEwfScKCiMgU2VuZCBub2RlIHRvIGVz
cGVjaWZpYyBkZXNrdG9wCmFsdCArIHNoaWZ0ICsgezEtOSwwfQoJYnNwYyBub
2RlIC1kICdeezEtOSwxMH0nCg==`

const POLYBAR_CONFIG string = `W2NvbG9yc10KYmFja2dyb3VuZCA9IC
MyODJBMkUKYmFja2dyb3VuZC1hbHQgPSAjMzczQjQxCmZvcmVncm91bmQgPSA
jQzVDOEM2CnByaW1hcnkgPSAjRjBDNjc0CnNlY29uZGFyeSA9ICM4QUJFQjcK
YWxlcnQgPSAjQTU0MjQyCmRpc2FibGVkID0gIzcwNzg4MAoKW2Jhci9leGFtc
GxlXQp3aWR0aCA9IDEwMCUKaGVpZ2h0ID0gMjRwdApyYWRpdXMgPSA2Cgo7IG
RwaSA9IDk2CgpiYWNrZ3JvdW5kID0gJHtjb2xvcnMuYmFja2dyb3VuZH0KZm9
yZWdyb3VuZCA9ICR7Y29sb3JzLmZvcmVncm91bmR9CgpsaW5lLXNpemUgPSAz
cHQKCmJvcmRlci1zaXplID0gNHB0CmJvcmRlci1jb2xvciA9ICMwMDAwMDAwM
AoKcGFkZGluZy1sZWZ0ID0gMApwYWRkaW5nLXJpZ2h0ID0gMQoKbW9kdWxlLW
1hcmdpbiA9IDEKCnNlcGFyYXRvciA9IHwKc2VwYXJhdG9yLWZvcmVncm91bmQ
gPSAke2NvbG9ycy5kaXNhYmxlZH0KCmZvbnQtMCA9IG1vbm9zcGFjZTsyCgpt
b2R1bGVzLWxlZnQgPSB4d29ya3NwYWNlcyB4d2luZG93Cm1vZHVsZXMtcmlna
HQgPSBmaWxlc3lzdGVtIHB1bHNlYXVkaW8geGtleWJvYXJkIG1lbW9yeSBjcH
Ugd2xhbiBldGggZGF0ZQoKY3Vyc29yLWNsaWNrID0gcG9pbnRlcgpjdXJzb3I
tc2Nyb2xsID0gbnMtcmVzaXplCgplbmFibGUtaXBjID0gdHJ1ZQoKOyB0cmF5
LXBvc2l0aW9uID0gcmlnaHQKCjsgd20tcmVzdGFjayA9IGdlbmVyaWMKOyB3b
S1yZXN0YWNrID0gYnNwd20KOyB3bS1yZXN0YWNrID0gaTMKCjsgb3ZlcnJpZG
UtcmVkaXJlY3QgPSB0cnVlCgpbbW9kdWxlL3h3b3Jrc3BhY2VzXQp0eXBlID0
gaW50ZXJuYWwveHdvcmtzcGFjZXMKCmxhYmVsLWFjdGl2ZSA9ICVuYW1lJQps
YWJlbC1hY3RpdmUtYmFja2dyb3VuZCA9ICR7Y29sb3JzLmJhY2tncm91bmQtY
Wx0fQpsYWJlbC1hY3RpdmUtdW5kZXJsaW5lPSAke2NvbG9ycy5wcmltYXJ5fQ
psYWJlbC1hY3RpdmUtcGFkZGluZyA9IDEKCmxhYmVsLW9jY3VwaWVkID0gJW5
hbWUlCmxhYmVsLW9jY3VwaWVkLXBhZGRpbmcgPSAxCgpsYWJlbC11cmdlbnQg
PSAlbmFtZSUKbGFiZWwtdXJnZW50LWJhY2tncm91bmQgPSAke2NvbG9ycy5hb
GVydH0KbGFiZWwtdXJnZW50LXBhZGRpbmcgPSAxCgpsYWJlbC1lbXB0eSA9IC
VuYW1lJQpsYWJlbC1lbXB0eS1mb3JlZ3JvdW5kID0gJHtjb2xvcnMuZGlzYWJ
sZWR9CmxhYmVsLWVtcHR5LXBhZGRpbmcgPSAxCgpbbW9kdWxlL3h3aW5kb3dd
CnR5cGUgPSBpbnRlcm5hbC94d2luZG93CmxhYmVsID0gJXRpdGxlOjA6NjA6L
i4uJQoKW21vZHVsZS9maWxlc3lzdGVtXQp0eXBlID0gaW50ZXJuYWwvZnMKaW
50ZXJ2YWwgPSAyNQoKbW91bnQtMCA9IC8KCmxhYmVsLW1vdW50ZWQgPSAle0Y
jRjBDNjc0fSVtb3VudHBvaW50JSV7Ri19ICVwZXJjZW50YWdlX3VzZWQlJQoK
bGFiZWwtdW5tb3VudGVkID0gJW1vdW50cG9pbnQlIG5vdCBtb3VudGVkCmxhY
mVsLXVubW91bnRlZC1mb3JlZ3JvdW5kID0gJHtjb2xvcnMuZGlzYWJsZWR9Cg
pbbW9kdWxlL3B1bHNlYXVkaW9dCnR5cGUgPSBpbnRlcm5hbC9wdWxzZWF1ZGl
vCgpmb3JtYXQtdm9sdW1lLXByZWZpeCA9ICJWT0wgIgpmb3JtYXQtdm9sdW1l
LXByZWZpeC1mb3JlZ3JvdW5kID0gJHtjb2xvcnMucHJpbWFyeX0KZm9ybWF0L
XZvbHVtZSA9IDxsYWJlbC12b2x1bWU+CgpsYWJlbC12b2x1bWUgPSAlcGVyY2
VudGFnZSUlCgpsYWJlbC1tdXRlZCA9IG11dGVkCmxhYmVsLW11dGVkLWZvcmV
ncm91bmQgPSAke2NvbG9ycy5kaXNhYmxlZH0KClttb2R1bGUveGtleWJvYXJk
XQp0eXBlID0gaW50ZXJuYWwveGtleWJvYXJkCmJsYWNrbGlzdC0wID0gbnVtI
GxvY2sKCmxhYmVsLWxheW91dCA9ICVsYXlvdXQlCmxhYmVsLWxheW91dC1mb3
JlZ3JvdW5kID0gJHtjb2xvcnMucHJpbWFyeX0KCmxhYmVsLWluZGljYXRvci1
wYWRkaW5nID0gMgpsYWJlbC1pbmRpY2F0b3ItbWFyZ2luID0gMQpsYWJlbC1p
bmRpY2F0b3ItZm9yZWdyb3VuZCA9ICR7Y29sb3JzLmJhY2tncm91bmR9CmxhY
mVsLWluZGljYXRvci1iYWNrZ3JvdW5kID0gJHtjb2xvcnMuc2Vjb25kYXJ5fQ
oKW21vZHVsZS9tZW1vcnldCnR5cGUgPSBpbnRlcm5hbC9tZW1vcnkKaW50ZXJ
2YWwgPSAyCmZvcm1hdC1wcmVmaXggPSAiUkFNICIKZm9ybWF0LXByZWZpeC1m
b3JlZ3JvdW5kID0gJHtjb2xvcnMucHJpbWFyeX0KbGFiZWwgPSAlcGVyY2Vud
GFnZV91c2VkOjIlJQoKW21vZHVsZS9jcHVdCnR5cGUgPSBpbnRlcm5hbC9jcH
UKaW50ZXJ2YWwgPSAyCmZvcm1hdC1wcmVmaXggPSAiQ1BVICIKZm9ybWF0LXB
yZWZpeC1mb3JlZ3JvdW5kID0gJHtjb2xvcnMucHJpbWFyeX0KbGFiZWwgPSAl
cGVyY2VudGFnZToyJSUKCltuZXR3b3JrLWJhc2VdCnR5cGUgPSBpbnRlcm5hb
C9uZXR3b3JrCmludGVydmFsID0gNQpmb3JtYXQtY29ubmVjdGVkID0gPGxhYm
VsLWNvbm5lY3RlZD4KZm9ybWF0LWRpc2Nvbm5lY3RlZCA9IDxsYWJlbC1kaXN
jb25uZWN0ZWQ+CmxhYmVsLWRpc2Nvbm5lY3RlZCA9ICV7RiNGMEM2NzR9JWlm
bmFtZSUle0YjNzA3ODgwfSBkaXNjb25uZWN0ZWQKClttb2R1bGUvd2xhbl0Ka
W5oZXJpdCA9IG5ldHdvcmstYmFzZQppbnRlcmZhY2UtdHlwZSA9IHdpcmVsZX
NzCmxhYmVsLWNvbm5lY3RlZCA9ICV7RiNGMEM2NzR9JWlmbmFtZSUle0YtfSA
lZXNzaWQlICVsb2NhbF9pcCUKClttb2R1bGUvZXRoXQppbmhlcml0ID0gbmV0
d29yay1iYXNlCmludGVyZmFjZS10eXBlID0gd2lyZWQKbGFiZWwtY29ubmVjd
GVkID0gJXtGI0YwQzY3NH0laWZuYW1lJSV7Ri19ICVsb2NhbF9pcCUKClttb2
R1bGUvZGF0ZV0KdHlwZSA9IGludGVybmFsL2RhdGUKaW50ZXJ2YWwgPSAxCgp
kYXRlID0gJUg6JU0KZGF0ZS1hbHQgPSAlWS0lbS0lZCAlSDolTTolUwoKbGFi
ZWwgPSAlZGF0ZSUKbGFiZWwtZm9yZWdyb3VuZCA9ICR7Y29sb3JzLnByaW1hc
nl9Cgpbc2V0dGluZ3NdCnNjcmVlbmNoYW5nZS1yZWxvYWQgPSB0cnVlCnBzZX
Vkby10cmFuc3BhcmVuY3kgPSB0cnVlCgo7IHZpbTpmdD1kb3NpbmkK`

const POLYBAR_LAUNCHER string = `IyEvYmluL2Jhc2gKCiMgVGVybWlu
YXRlIGFscmVhZHkgcnVubmluZyBiYXIgaW5zdGFuY2VzCmtpbGxhbGwgLXEgc
G9seWJhcgojIElmIGFsbCB5b3VyIGJhcnMgaGF2ZSBpcGMgZW5hYmxlZCwgeW
91IGNhbiBhbHNvIHVzZSAKIyBwb2x5YmFyLW1zZyBjbWQgcXVpdAoKIyBMYXV
uY2ggUG9seWJhciwgdXNpbmcgZGVmYXVsdCBjb25maWcgbG9jYXRpb24gfi8u
Y29uZmlnL3BvbHliYXIvY29uZmlnLmluaQpwb2x5YmFyIGV4YW1wbGUgMj4mM
SB8IHRlZSAtYSAvdG1wL3BvbHliYXIubG9nICYgZGlzb3duCgplY2hvICJQb2
x5YmFyIGxhdW5jaGVkLi4uIgo=`

const NITROGEN_CONFIG string = `W2dlb21ldHJ5XQpwb3N4PTEyCnBvc
3k9NTQKc2l6ZXg9MTg5NApzaXpleT0xMDEyCgpbbml0cm9nZW5dCnZpZXc9aW
NvbgpyZWN1cnNlPXRydWUKc29ydD1hbHBoYQppY29uX2NhcHM9ZmFsc2UKZGl
ycz0vdXNyL3NoYXJlL2JhY2tncm91bmRzL2FyY2hsaW51eDs=`

const NITROGEN_SAVE string = `W3hpbl8tMV0KZmlsZT0vdXNyL3NoYXJ
lL2JhY2tncm91bmRzL2FyY2hsaW51eC9zcGxpdC5wbmcKbW9kZT00CmJnY29s
b3I9IzAwMDAwMA==`

const XFCE4_TERMINAL_CONFIG string = `W0NvbmZpZ3VyYXRpb25dCk1
pc2NBbHdheXNTaG93VGFicz1GQUxTRQpNaXNjQmVsbD1GQUxTRQpNaXNjQmVs
bFVyZ2VudD1GQUxTRQpNaXNjQm9yZGVyc0RlZmF1bHQ9RkFMU0UKTWlzY0N1c
nNvckJsaW5rcz1GQUxTRQpNaXNjQ3Vyc29yU2hhcGU9VEVSTUlOQUxfQ1VSU0
9SX1NIQVBFX0lCRUFNCk1pc2NEZWZhdWx0R2VvbWV0cnk9ODB4MjQKTWlzY0l
uaGVyaXRHZW9tZXRyeT1GQUxTRQpNaXNjTWVudWJhckRlZmF1bHQ9RkFMU0UK
TWlzY01vdXNlQXV0b2hpZGU9RkFMU0UKTWlzY01vdXNlV2hlZWxab29tPVRSV
UUKTWlzY1Rvb2xiYXJEZWZhdWx0PUZBTFNFCk1pc2NDb25maXJtQ2xvc2U9VF
JVRQpNaXNjQ3ljbGVUYWJzPVRSVUUKTWlzY1RhYkNsb3NlQnV0dG9ucz1UUlV
FCk1pc2NUYWJDbG9zZU1pZGRsZUNsaWNrPVRSVUUKTWlzY1RhYlBvc2l0aW9u
PUdUS19QT1NfVE9QCk1pc2NIaWdobGlnaHRVcmxzPVRSVUUKTWlzY01pZGRsZ
UNsaWNrT3BlbnNVcmk9RkFMU0UKTWlzY0NvcHlPblNlbGVjdD1GQUxTRQpNaX
NjU2hvd1JlbGF1bmNoRGlhbG9nPVRSVUUKTWlzY1Jld3JhcE9uUmVzaXplPVR
SVUUKTWlzY1VzZVNoaWZ0QXJyb3dzVG9TY3JvbGw9RkFMU0UKTWlzY1NsaW1U
YWJzPUZBTFNFCk1pc2NOZXdUYWJBZGphY2VudD1GQUxTRQpNaXNjU2VhcmNoR
GlhbG9nT3BhY2l0eT0xMDAKTWlzY1Nob3dVbnNhZmVQYXN0ZURpYWxvZz1UUl
VFCk1pc2NSaWdodENsaWNrQWN0aW9uPVRFUk1JTkFMX1JJR0hUX0NMSUNLX0F
DVElPTl9DT05URVhUX01FTlUKU2Nyb2xsaW5nQmFyPVRFUk1JTkFMX1NDUk9M
TEJBUl9OT05FCkJhY2tncm91bmRNb2RlPVRFUk1JTkFMX0JBQ0tHUk9VTkRfV
FJBTlNQQVJFTlQK`

const BASHRC string = `WyAteiAiJFBTMSIgXSAmJiByZXR1cm4KCkhJU1
RDT05UUk9MPWlnbm9yZWR1cHM6aWdub3Jlc3BhY2UKCnNob3B0IC1zIGhpc3R
hcHBlbmQKCkhJU1RTSVpFPTEwMDAKSElTVEZJTEVTSVpFPTIwMDAKCnNob3B0
IC1zIGNoZWNrd2luc2l6ZQoKWyAteCAvdXNyL2Jpbi9sZXNzcGlwZSBdICYmI
GV2YWwgIiQoU0hFTEw9L2Jpbi9zaCBsZXNzcGlwZSkiCgppZiBbIC16ICIkZG
ViaWFuX2Nocm9vdCIgXSAmJiBbIC1yIC9ldGMvZGViaWFuX2Nocm9vdCBdOyB
0aGVuCiAgICBkZWJpYW5fY2hyb290PSQoY2F0IC9ldGMvZGViaWFuX2Nocm9v
dCkKZmkKCmNhc2UgIiRURVJNIiBpbgogICAgeHRlcm0tY29sb3IpIGNvbG9yX
3Byb21wdD15ZXM7Owplc2FjCgppZiBbIC1uICIkZm9yY2VfY29sb3JfcHJvbX
B0IiBdOyB0aGVuCiAgICBpZiBbIC14IC91c3IvYmluL3RwdXQgXSAmJiB0cHV
0IHNldGFmIDEgPiYvZGV2L251bGw7IHRoZW4KCSMgV2UgaGF2ZSBjb2xvciBz
dXBwb3J0OyBhc3N1bWUgaXQncyBjb21wbGlhbnQgd2l0aCBFY21hLTQ4CgkjI
ChJU08vSUVDLTY0MjkpLiAoTGFjayBvZiBzdWNoIHN1cHBvcnQgaXMgZXh0cm
VtZWx5IHJhcmUsIGFuZCBzdWNoCgkjIGEgY2FzZSB3b3VsZCB0ZW5kIHRvIHN
1cHBvcnQgc2V0ZiByYXRoZXIgdGhhbiBzZXRhZi4pCgljb2xvcl9wcm9tcHQ9
eWVzCiAgICBlbHNlCgljb2xvcl9wcm9tcHQ9CiAgICBmaQpmaQoKaWYgWyAiJ
GNvbG9yX3Byb21wdCIgPSB5ZXMgXTsgdGhlbgogICAgUFMxPSJcW1wwMzNbMD
szMW1cXVwzNDJcMjI0XDIxNFwzNDJcMjI0XDIwMFwkKFtbIFwkPyAhPSAwIF1
dICYmIGVjaG8gXCJbXFtcMDMzWzA7MzFtXF1cMzQyXDIzNFwyMjdcW1wwMzNb
MDszN21cXV1cMzQyXDIyNFwyMDBcIilbJChpZiBbWyAke0VVSUR9ID09IDAgX
V07IHRoZW4gZWNobyAnXFtcMDMzWzAxOzMxbVxdcm9vdFxbXDAzM1swMTszM2
1cXUBcW1wwMzNbMDE7OTZtXF1caCc7IGVsc2UgZWNobyAnXFtcMDMzWzA7Mzl
tXF1cdVxbXDAzM1swMTszM21cXUBcW1wwMzNbMDE7OTZtXF1caCc7IGZpKVxb
XDAzM1swOzMxbVxdXVwzNDJcMjI0XDIwMFtcW1wwMzNbMDszMm1cXVx3XFtcM
DMzWzA7MzFtXF1dXG5cW1wwMzNbMDszMW1cXVwzNDJcMjI0XDIyNFwzNDJcMj
I0XDIwMFwzNDJcMjI0XDIwMFwzNDJcMjI1XDI3NCBcW1wwMzNbMG1cXVxbXGV
bMDE7MzNtXF1cXCRcW1xlWzBtXF0iCmVsc2UKICAgIFBTMT0iXFtcMDMzWzA7
MzFtXF1cMzQyXDIyNFwyMTRcMzQyXDIyNFwyMDBcJChbWyBcJD8gIT0gMCBdX
SAmJiBlY2hvIFwiW1xbXDAzM1swOzMxbVxdXDM0MlwyMzRcMjI3XFtcMDMzWz
A7MzdtXF1dXDM0MlwyMjRcMjAwXCIpWyQoaWYgW1sgJHtFVUlEfSA9PSAwIF1
dOyB0aGVuIGVjaG8gJ1xbXDAzM1swMTszMW1cXXJvb3RcW1wwMzNbMDE7MzNt
XF1AXFtcMDMzWzAxOzk2bVxdXGgnOyBlbHNlIGVjaG8gJ1xbXDAzM1swOzM5b
VxdXHVcW1wwMzNbMDE7MzNtXF1AXFtcMDMzWzAxOzk2bVxdXGgnOyBmaSlcW1
wwMzNbMDszMW1cXV1cMzQyXDIyNFwyMDBbXFtcMDMzWzA7MzJtXF1cd1xbXDA
zM1swOzMxbVxdXVxuXFtcMDMzWzA7MzFtXF1cMzQyXDIyNFwyMjRcMzQyXDIy
NFwyMDBcMzQyXDIyNFwyMDBcMzQyXDIyNVwyNzQgXFtcMDMzWzBtXF1cW1xlW
zAxOzMzbVxdXFwkXFtcZVswbVxdIgpmaQp1bnNldCBjb2xvcl9wcm9tcHQgZm
9yY2VfY29sb3JfcHJvbXB0CgpjYXNlICIkVEVSTSIgaW4KeHRlcm0qfHJ4dnQ
qKQogICAgUFMxPSJcW1wwMzNbMDszMW1cXVwzNDJcMjI0XDIxNFwzNDJcMjI0
XDIwMFwkKFtbIFwkPyAhPSAwIF1dICYmIGVjaG8gXCJbXFtcMDMzWzA7MzFtX
F1cMzQyXDIzNFwyMjdcW1wwMzNbMDszN21cXV1cMzQyXDIyNFwyMDBcIilbJC
hpZiBbWyAke0VVSUR9ID09IDAgXV07IHRoZW4gZWNobyAnXFtcMDMzWzAxOzM
xbVxdcm9vdFxbXDAzM1swMTszM21cXUBcW1wwMzNbMDE7OTZtXF1caCc7IGVs
c2UgZWNobyAnXFtcMDMzWzA7MzltXF1cdVxbXDAzM1swMTszM21cXUBcW1wwM
zNbMDE7OTZtXF1caCc7IGZpKVxbXDAzM1swOzMxbVxdXVwzNDJcMjI0XDIwMF
tcW1wwMzNbMDszMm1cXVx3XFtcMDMzWzA7MzFtXF1dXG5cW1wwMzNbMDszMW1
cXVwzNDJcMjI0XDIyNFwzNDJcMjI0XDIwMFwzNDJcMjI0XDIwMFwzNDJcMjI1
XDI3NCBcW1wwMzNbMG1cXVxbXGVbMDE7MzNtXF1cXCRcW1xlWzBtXF0iCiAgI
CAjUFMxPSJcW1xlXTA7JHtkZWJpYW5fY2hyb290OisoJGRlYmlhbl9jaHJvb3
QpfVx1QFxoOiBcd1xhXF0kUFMxIgogICAgOzsKKikKICAgIDs7CmVzYWMKCml
mIFsgLXggL3Vzci9iaW4vZGlyY29sb3JzIF07IHRoZW4KICAgIHRlc3QgLXIg
fi8uZGlyY29sb3JzICYmIGV2YWwgIiQoZGlyY29sb3JzIC1iIH4vLmRpcmNvb
G9ycykiIHx8IGV2YWwgIiQoZGlyY29sb3JzIC1iKSIKICAgIGFsaWFzIGxzPS
dscyAtLWNvbG9yPWF1dG8nCiAgICAjYWxpYXMgZGlyPSdkaXIgLS1jb2xvcj1
hdXRvJwogICAgI2FsaWFzIHZkaXI9J3ZkaXIgLS1jb2xvcj1hdXRvJwoKICAg
IGFsaWFzIGdyZXA9J2dyZXAgLS1jb2xvcj1hdXRvJwogICAgYWxpYXMgZmdyZ
XA9J2ZncmVwIC0tY29sb3I9YXV0bycKICAgIGFsaWFzIGVncmVwPSdlZ3JlcC
AtLWNvbG9yPWF1dG8nCmZpCgphbGlhcyBsbD0nbHMgLWFsRicKYWxpYXMgbGE
9J2xzIC1BJwphbGlhcyBsPSdscyAtQ0YnCgphbGlhcyBhbGVydD0nbm90aWZ5
LXNlbmQgLS11cmdlbmN5PWxvdyAtaSAiJChbICQ/ID0gMCBdICYmIGVjaG8gd
GVybWluYWwgfHwgZWNobyBlcnJvcikiICIkKGhpc3Rvcnl8dGFpbCAtbjF8c2
VkIC1lICdcJydzL15ccypbMC05XVwrXHMqLy87cy9bOyZ8XVxzKmFsZXJ0JC8
vJ1wnJykiJwoKaWYgWyAtZiB+Ly5iYXNoX2FsaWFzZXMgXTsgdGhlbgogICAg
LiB+Ly5iYXNoX2FsaWFzZXMKZmkKCmlmIFsgLWYgL2V0Yy9iYXNoX2NvbXBsZ
XRpb24gXSAmJiAhIHNob3B0IC1vcSBwb3NpeDsgdGhlbgogICAgLiAvZXRjL2
Jhc2hfY29tcGxldGlvbgpmaQoKYWxpYXMgY2xzPSdjbGVhcicKYWxpYXMgcHk
9J3B5dGhvbjMnCg==`
