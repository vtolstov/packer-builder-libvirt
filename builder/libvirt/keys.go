package libvirt

const (
	KEY_RESERVED                 = 0
	KEY_ESC                      = 1
	KEY_1                        = 2
	KEY_2                        = 3
	KEY_3                        = 4
	KEY_4                        = 5
	KEY_5                        = 6
	KEY_6                        = 7
	KEY_7                        = 8
	KEY_8                        = 9
	KEY_9                        = 10
	KEY_0                        = 11
	KEY_MINUS                    = 12
	KEY_EQUAL                    = 13
	KEY_BACKSPACE                = 14
	KEY_TAB                      = 15
	KEY_Q                        = 16
	KEY_W                        = 17
	KEY_E                        = 18
	KEY_R                        = 19
	KEY_T                        = 20
	KEY_Y                        = 21
	KEY_U                        = 22
	KEY_I                        = 23
	KEY_O                        = 24
	KEY_P                        = 25
	KEY_LEFTBRACE                = 26
	KEY_RIGHTBRACE               = 27
	KEY_ENTER                    = 28
	KEY_LEFTCTRL                 = 29
	KEY_A                        = 30
	KEY_S                        = 31
	KEY_D                        = 32
	KEY_F                        = 33
	KEY_G                        = 34
	KEY_H                        = 35
	KEY_J                        = 36
	KEY_K                        = 37
	KEY_L                        = 38
	KEY_SEMICOLON                = 39
	KEY_APOSTROPHE               = 40
	KEY_GRAVE                    = 41
	KEY_LEFTSHIFT                = 42
	KEY_BACKSLASH                = 43
	KEY_Z                        = 44
	KEY_X                        = 45
	KEY_C                        = 46
	KEY_V                        = 47
	KEY_B                        = 48
	KEY_N                        = 49
	KEY_M                        = 50
	KEY_COMMA                    = 51
	KEY_DOT                      = 52
	KEY_SLASH                    = 53
	KEY_RIGHTSHIFT               = 54
	KEY_KPASTERISK               = 55
	KEY_LEFTALT                  = 56
	KEY_SPACE                    = 57
	KEY_CAPSLOCK                 = 58
	KEY_F1                       = 59
	KEY_F2                       = 60
	KEY_F3                       = 61
	KEY_F4                       = 62
	KEY_F5                       = 63
	KEY_F6                       = 64
	KEY_F7                       = 65
	KEY_F8                       = 66
	KEY_F9                       = 67
	KEY_F10                      = 68
	KEY_NUMLOCK                  = 69
	KEY_SCROLLLOCK               = 70
	KEY_KP7                      = 71
	KEY_KP8                      = 72
	KEY_KP9                      = 73
	KEY_KPMINUS                  = 74
	KEY_KP4                      = 75
	KEY_KP5                      = 76
	KEY_KP6                      = 77
	KEY_KPPLUS                   = 78
	KEY_KP1                      = 79
	KEY_KP2                      = 80
	KEY_KP3                      = 81
	KEY_KP0                      = 82
	KEY_KPDOT                    = 83
	KEY_ZENKAKUHANKAKU           = 85
	KEY_102ND                    = 86
	KEY_F11                      = 87
	KEY_F12                      = 88
	KEY_RO                       = 89
	KEY_KATAKANA                 = 90
	KEY_HIRAGANA                 = 91
	KEY_HENKAN                   = 92
	KEY_KATAKANAHIRAGANA         = 93
	KEY_MUHENKAN                 = 94
	KEY_KPJPCOMMA                = 95
	KEY_KPENTER                  = 96
	KEY_RIGHTCTRL                = 97
	KEY_KPSLASH                  = 98
	KEY_SYSRQ                    = 99
	KEY_RIGHTALT                 = 100
	KEY_LINEFEED                 = 101
	KEY_HOME                     = 102
	KEY_UP                       = 103
	KEY_PAGEUP                   = 104
	KEY_LEFT                     = 105
	KEY_RIGHT                    = 106
	KEY_END                      = 107
	KEY_DOWN                     = 108
	KEY_PAGEDOWN                 = 109
	KEY_INSERT                   = 110
	KEY_DELETE                   = 111
	KEY_MACRO                    = 112
	KEY_MUTE                     = 113
	KEY_VOLUMEDOWN               = 114
	KEY_VOLUMEUP                 = 115
	KEY_POWER                    = 116
	KEY_KPEQUAL                  = 117
	KEY_KPPLUSMINUS              = 118
	KEY_PAUSE                    = 119
	KEY_SCALE                    = 120
	KEY_KPCOMMA                  = 121
	KEY_HANGEUL                  = 122
	KEY_HANGUEL                  = KEY_HANGEUL
	KEY_HANJA                    = 123
	KEY_YEN                      = 124
	KEY_LEFTMETA                 = 125
	KEY_RIGHTMETA                = 126
	KEY_COMPOSE                  = 127
	KEY_STOP                     = 128
	KEY_AGAIN                    = 129
	KEY_PROPS                    = 130
	KEY_UNDO                     = 131
	KEY_FRONT                    = 132
	KEY_COPY                     = 133
	KEY_OPEN                     = 134
	KEY_PASTE                    = 135
	KEY_FIND                     = 136
	KEY_CUT                      = 137
	KEY_HELP                     = 138
	KEY_MENU                     = 139
	KEY_CALC                     = 140
	KEY_SETUP                    = 141
	KEY_SLEEP                    = 142
	KEY_WAKEUP                   = 143
	KEY_FILE                     = 144
	KEY_SENDFILE                 = 145
	KEY_DELETEFILE               = 146
	KEY_XFER                     = 147
	KEY_PROG1                    = 148
	KEY_PROG2                    = 149
	KEY_WWW                      = 150
	KEY_MSDOS                    = 151
	KEY_COFFEE                   = 152
	KEY_SCREENLOCK               = KEY_COFFEE
	KEY_DIRECTION                = 153
	KEY_CYCLEWINDOWS             = 154
	KEY_MAIL                     = 155
	KEY_BOOKMARKS                = 156
	KEY_COMPUTER                 = 157
	KEY_BACK                     = 158
	KEY_FORWARD                  = 159
	KEY_CLOSECD                  = 160
	KEY_EJECTCD                  = 161
	KEY_EJECTCLOSECD             = 162
	KEY_NEXTSONG                 = 163
	KEY_PLAYPAUSE                = 164
	KEY_PREVIOUSSONG             = 165
	KEY_STOPCD                   = 166
	KEY_RECORD                   = 167
	KEY_REWIND                   = 168
	KEY_PHONE                    = 169
	KEY_ISO                      = 170
	KEY_CONFIG                   = 171
	KEY_HOMEPAGE                 = 172
	KEY_REFRESH                  = 173
	KEY_EXIT                     = 174
	KEY_MOVE                     = 175
	KEY_EDIT                     = 176
	KEY_SCROLLUP                 = 177
	KEY_SCROLLDOWN               = 178
	KEY_KPLEFTPAREN              = 179
	KEY_KPRIGHTPAREN             = 180
	KEY_NEW                      = 181
	KEY_REDO                     = 182
	KEY_F13                      = 183
	KEY_F14                      = 184
	KEY_F15                      = 185
	KEY_F16                      = 186
	KEY_F17                      = 187
	KEY_F18                      = 188
	KEY_F19                      = 189
	KEY_F20                      = 190
	KEY_F21                      = 191
	KEY_F22                      = 192
	KEY_F23                      = 193
	KEY_F24                      = 194
	KEY_PLAYCD                   = 200
	KEY_PAUSECD                  = 201
	KEY_PROG3                    = 202
	KEY_PROG4                    = 203
	KEY_DASHBOARD                = 204
	KEY_SUSPEND                  = 205
	KEY_CLOSE                    = 206
	KEY_PLAY                     = 207
	KEY_FASTFORWARD              = 208
	KEY_BASSBOOST                = 209
	KEY_PRINT                    = 210
	KEY_HP                       = 211
	KEY_CAMERA                   = 212
	KEY_SOUND                    = 213
	KEY_QUESTION                 = 214
	KEY_EMAIL                    = 215
	KEY_CHAT                     = 216
	KEY_SEARCH                   = 217
	KEY_CONNECT                  = 218
	KEY_FINANCE                  = 219
	KEY_SPORT                    = 220
	KEY_SHOP                     = 221
	KEY_ALTERASE                 = 222
	KEY_CANCEL                   = 223
	KEY_BRIGHTNESSDOWN           = 224
	KEY_BRIGHTNESSUP             = 225
	KEY_MEDIA                    = 226
	KEY_SWITCHVIDEOMODE          = 227
	KEY_KBDILLUMTOGGLE           = 228
	KEY_KBDILLUMDOWN             = 229
	KEY_KBDILLUMUP               = 230
	KEY_SEND                     = 231
	KEY_REPLY                    = 232
	KEY_FORWARDMAIL              = 233
	KEY_SAVE                     = 234
	KEY_DOCUMENTS                = 235
	KEY_BATTERY                  = 236
	KEY_BLUETOOTH                = 237
	KEY_WLAN                     = 238
	KEY_UWB                      = 239
	KEY_UNKNOWN                  = 240
	KEY_VIDEO_NEXT               = 241
	KEY_VIDEO_PREV               = 242
	KEY_BRIGHTNESS_CYCLE         = 243
	KEY_BRIGHTNESS_AUTO          = 244
	KEY_BRIGHTNESS_ZERO          = KEY_BRIGHTNESS_AUTO
	KEY_DISPLAY_OFF              = 245
	KEY_WWAN                     = 246
	KEY_WIMAX                    = KEY_WWAN
	KEY_RFKILL                   = 247
	KEY_MICMUTE                  = 248
	KEY_OK                       = 0x160
	KEY_SELECT                   = 0x161
	KEY_GOTO                     = 0x162
	KEY_CLEAR                    = 0x163
	KEY_POWER2                   = 0x164
	KEY_OPTION                   = 0x165
	KEY_INFO                     = 0x166
	KEY_TIME                     = 0x167
	KEY_VENDOR                   = 0x168
	KEY_ARCHIVE                  = 0x169
	KEY_PROGRAM                  = 0x16a
	KEY_CHANNEL                  = 0x16b
	KEY_FAVORITES                = 0x16c
	KEY_EPG                      = 0x16d
	KEY_PVR                      = 0x16e
	KEY_MHP                      = 0x16f
	KEY_LANGUAGE                 = 0x170
	KEY_TITLE                    = 0x171
	KEY_SUBTITLE                 = 0x172
	KEY_ANGLE                    = 0x173
	KEY_ZOOM                     = 0x174
	KEY_MODE                     = 0x175
	KEY_KEYBOARD                 = 0x176
	KEY_SCREEN                   = 0x177
	KEY_PC                       = 0x178
	KEY_TV                       = 0x179
	KEY_TV2                      = 0x17a
	KEY_VCR                      = 0x17b
	KEY_VCR2                     = 0x17c
	KEY_SAT                      = 0x17d
	KEY_SAT2                     = 0x17e
	KEY_CD                       = 0x17f
	KEY_TAPE                     = 0x180
	KEY_RADIO                    = 0x181
	KEY_TUNER                    = 0x182
	KEY_PLAYER                   = 0x183
	KEY_TEXT                     = 0x184
	KEY_DVD                      = 0x185
	KEY_AUX                      = 0x186
	KEY_MP3                      = 0x187
	KEY_AUDIO                    = 0x188
	KEY_VIDEO                    = 0x189
	KEY_DIRECTORY                = 0x18a
	KEY_LIST                     = 0x18b
	KEY_MEMO                     = 0x18c
	KEY_CALENDAR                 = 0x18d
	KEY_RED                      = 0x18e
	KEY_GREEN                    = 0x18f
	KEY_YELLOW                   = 0x190
	KEY_BLUE                     = 0x191
	KEY_CHANNELUP                = 0x192
	KEY_CHANNELDOWN              = 0x193
	KEY_FIRST                    = 0x194
	KEY_LAST                     = 0x195
	KEY_AB                       = 0x196
	KEY_NEXT                     = 0x197
	KEY_RESTART                  = 0x198
	KEY_SLOW                     = 0x199
	KEY_SHUFFLE                  = 0x19a
	KEY_BREAK                    = 0x19b
	KEY_PREVIOUS                 = 0x19c
	KEY_DIGITS                   = 0x19d
	KEY_TEEN                     = 0x19e
	KEY_TWEN                     = 0x19f
	KEY_VIDEOPHONE               = 0x1a0
	KEY_GAMES                    = 0x1a1
	KEY_ZOOMIN                   = 0x1a2
	KEY_ZOOMOUT                  = 0x1a3
	KEY_ZOOMRESET                = 0x1a4
	KEY_WORDPROCESSOR            = 0x1a5
	KEY_EDITOR                   = 0x1a6
	KEY_SPREADSHEET              = 0x1a7
	KEY_GRAPHICSEDITOR           = 0x1a8
	KEY_PRESENTATION             = 0x1a9
	KEY_DATABASE                 = 0x1aa
	KEY_NEWS                     = 0x1ab
	KEY_VOICEMAIL                = 0x1ac
	KEY_ADDRESSBOOK              = 0x1ad
	KEY_MESSENGER                = 0x1ae
	KEY_DISPLAYTOGGLE            = 0x1af
	KEY_BRIGHTNESS_TOGGLE        = KEY_DISPLAYTOGGLE
	KEY_SPELLCHECK               = 0x1b0
	KEY_LOGOFF                   = 0x1b1
	KEY_DOLLAR                   = 0x1b2
	KEY_EURO                     = 0x1b3
	KEY_FRAMEBACK                = 0x1b4
	KEY_FRAMEFORWARD             = 0x1b5
	KEY_CONTEXT_MENU             = 0x1b6
	KEY_MEDIA_REPEAT             = 0x1b7
	KEY_10CHANNELSUP             = 0x1b8
	KEY_10CHANNELSDOWN           = 0x1b9
	KEY_IMAGES                   = 0x1ba
	KEY_DEL_EOL                  = 0x1c0
	KEY_DEL_EOS                  = 0x1c1
	KEY_INS_LINE                 = 0x1c2
	KEY_DEL_LINE                 = 0x1c3
	KEY_FN                       = 0x1d0
	KEY_FN_ESC                   = 0x1d1
	KEY_FN_F1                    = 0x1d2
	KEY_FN_F2                    = 0x1d3
	KEY_FN_F3                    = 0x1d4
	KEY_FN_F4                    = 0x1d5
	KEY_FN_F5                    = 0x1d6
	KEY_FN_F6                    = 0x1d7
	KEY_FN_F7                    = 0x1d8
	KEY_FN_F8                    = 0x1d9
	KEY_FN_F9                    = 0x1da
	KEY_FN_F10                   = 0x1db
	KEY_FN_F11                   = 0x1dc
	KEY_FN_F12                   = 0x1dd
	KEY_FN_1                     = 0x1de
	KEY_FN_2                     = 0x1df
	KEY_FN_D                     = 0x1e0
	KEY_FN_E                     = 0x1e1
	KEY_FN_F                     = 0x1e2
	KEY_FN_S                     = 0x1e3
	KEY_FN_B                     = 0x1e4
	KEY_BRL_DOT1                 = 0x1f1
	KEY_BRL_DOT2                 = 0x1f2
	KEY_BRL_DOT3                 = 0x1f3
	KEY_BRL_DOT4                 = 0x1f4
	KEY_BRL_DOT5                 = 0x1f5
	KEY_BRL_DOT6                 = 0x1f6
	KEY_BRL_DOT7                 = 0x1f7
	KEY_BRL_DOT8                 = 0x1f8
	KEY_BRL_DOT9                 = 0x1f9
	KEY_BRL_DOT10                = 0x1fa
	KEY_NUMERIC_0                = 0x200
	KEY_NUMERIC_1                = 0x201
	KEY_NUMERIC_2                = 0x202
	KEY_NUMERIC_3                = 0x203
	KEY_NUMERIC_4                = 0x204
	KEY_NUMERIC_5                = 0x205
	KEY_NUMERIC_6                = 0x206
	KEY_NUMERIC_7                = 0x207
	KEY_NUMERIC_8                = 0x208
	KEY_NUMERIC_9                = 0x209
	KEY_NUMERIC_STAR             = 0x20a
	KEY_NUMERIC_POUND            = 0x20b
	KEY_CAMERA_FOCUS             = 0x210
	KEY_WPS_BUTTON               = 0x211
	KEY_TOUCHPAD_TOGGLE          = 0x212
	KEY_TOUCHPAD_ON              = 0x213
	KEY_TOUCHPAD_OFF             = 0x214
	KEY_CAMERA_ZOOMIN            = 0x215
	KEY_CAMERA_ZOOMOUT           = 0x216
	KEY_CAMERA_UP                = 0x217
	KEY_CAMERA_DOWN              = 0x218
	KEY_CAMERA_LEFT              = 0x219
	KEY_CAMERA_RIGHT             = 0x21a
	KEY_ATTENDANT_ON             = 0x21b
	KEY_ATTENDANT_OFF            = 0x21c
	KEY_ATTENDANT_TOGGLE         = 0x21d
	KEY_LIGHTS_TOGGLE            = 0x21e
	KEY_ALS_TOGGLE               = 0x230
	KEY_BUTTONCONFIG             = 0x240
	KEY_TASKMANAGER              = 0x241
	KEY_JOURNAL                  = 0x242
	KEY_CONTROLPANEL             = 0x243
	KEY_APPSELECT                = 0x244
	KEY_SCREENSAVER              = 0x245
	KEY_VOICECOMMAND             = 0x246
	KEY_BRIGHTNESS_MIN           = 0x250
	KEY_BRIGHTNESS_MAX           = 0x251
	KEY_KBDINPUTASSIST_PREV      = 0x260
	KEY_KBDINPUTASSIST_NEXT      = 0x261
	KEY_KBDINPUTASSIST_PREVGROUP = 0x262
	KEY_KBDINPUTASSIST_NEXTGROUP = 0x263
	KEY_KBDINPUTASSIST_ACCEPT    = 0x264
	KEY_KBDINPUTASSIST_CANCEL    = 0x265
	KEY_MIN_INTERESTING          = KEY_MUTE
	KEY_MAX                      = 0x2ff
	KEY_CNT                      = (KEY_MAX + 1)
)

var ecodes = map[string]uint{
	"<reserved>":                   KEY_RESERVED,
	"<esc>":                        KEY_ESC,
	"1":                            KEY_1,
	"2":                            KEY_2,
	"3":                            KEY_3,
	"4":                            KEY_4,
	"5":                            KEY_5,
	"6":                            KEY_6,
	"7":                            KEY_7,
	"8":                            KEY_8,
	"9":                            KEY_9,
	"0":                            KEY_0,
	"-":                            KEY_MINUS,
	"=":                            KEY_EQUAL,
	"<bs>":                         KEY_BACKSPACE,
	"<backspace>":                  KEY_BACKSPACE,
	"<tab>":                        KEY_TAB,
	"q":                            KEY_Q,
	"w":                            KEY_W,
	"e":                            KEY_E,
	"r":                            KEY_R,
	"t":                            KEY_T,
	"y":                            KEY_Y,
	"u":                            KEY_U,
	"i":                            KEY_I,
	"o":                            KEY_O,
	"p":                            KEY_P,
	"{":                            KEY_LEFTBRACE,
	"}":                            KEY_RIGHTBRACE,
	"<enter>":                      KEY_ENTER,
	"<lctrl>":                      KEY_LEFTCTRL,
	"a":                            KEY_A,
	"s":                            KEY_S,
	"d":                            KEY_D,
	"f":                            KEY_F,
	"g":                            KEY_G,
	"h":                            KEY_H,
	"j":                            KEY_J,
	"k":                            KEY_K,
	"l":                            KEY_L,
	";":                            KEY_SEMICOLON,
	"'":                            KEY_APOSTROPHE,
	"`":                            KEY_GRAVE,
	"<lshift>":                     KEY_LEFTSHIFT,
	"\\":                           KEY_BACKSLASH,
	"z":                            KEY_Z,
	"x":                            KEY_X,
	"c":                            KEY_C,
	"v":                            KEY_V,
	"b":                            KEY_B,
	"n":                            KEY_N,
	"m":                            KEY_M,
	",":                            KEY_COMMA,
	".":                            KEY_DOT,
	"/":                            KEY_SLASH,
	"<rshift>":                     KEY_RIGHTSHIFT,
	"KEY_KPASTERISK":               KEY_KPASTERISK,
	"<lalt>":                       KEY_LEFTALT,
	" ":                            KEY_SPACE,
	"<caps>":                       KEY_CAPSLOCK,
	"<f1>":                         KEY_F1,
	"<f2>":                         KEY_F2,
	"<f3>":                         KEY_F3,
	"<f4>":                         KEY_F4,
	"<f5>":                         KEY_F5,
	"<f6>":                         KEY_F6,
	"<f7>":                         KEY_F7,
	"<f8>":                         KEY_F8,
	"<f9>":                         KEY_F9,
	"<f10>":                        KEY_F10,
	"<numlock>":                    KEY_NUMLOCK,
	"<scrolllock>":                 KEY_SCROLLLOCK,
	"KEY_KP7":                      KEY_KP7,
	"KEY_KP8":                      KEY_KP8,
	"KEY_KP9":                      KEY_KP9,
	"KEY_KPMINUS":                  KEY_KPMINUS,
	"KEY_KP4":                      KEY_KP4,
	"KEY_KP5":                      KEY_KP5,
	"KEY_KP6":                      KEY_KP6,
	"KEY_KPPLUS":                   KEY_KPPLUS,
	"KEY_KP1":                      KEY_KP1,
	"KEY_KP2":                      KEY_KP2,
	"KEY_KP3":                      KEY_KP3,
	"KEY_KP0":                      KEY_KP0,
	"KEY_KPDOT":                    KEY_KPDOT,
	"KEY_ZENKAKUHANKAKU":           KEY_ZENKAKUHANKAKU,
	"KEY_102ND":                    KEY_102ND,
	"<f11>":                        KEY_F11,
	"<f12>":                        KEY_F12,
	"KEY_RO":                       KEY_RO,
	"KEY_KATAKANA":                 KEY_KATAKANA,
	"KEY_HIRAGANA":                 KEY_HIRAGANA,
	"KEY_HENKAN":                   KEY_HENKAN,
	"KEY_KATAKANAHIRAGANA":         KEY_KATAKANAHIRAGANA,
	"KEY_MUHENKAN":                 KEY_MUHENKAN,
	"KEY_KPJPCOMMA":                KEY_KPJPCOMMA,
	"KEY_KPENTER":                  KEY_KPENTER,
	"KEY_RIGHTCTRL":                KEY_RIGHTCTRL,
	"KEY_KPSLASH":                  KEY_KPSLASH,
	"<sysrq>":                      KEY_SYSRQ,
	"<ralt>":                       KEY_RIGHTALT,
	"KEY_LINEFEED":                 KEY_LINEFEED,
	"<home>":                       KEY_HOME,
	"<up>":                         KEY_UP,
	"<pageup>":                     KEY_PAGEUP,
	"<left>":                       KEY_LEFT,
	"<right>":                      KEY_RIGHT,
	"<end>":                        KEY_END,
	"<down>":                       KEY_DOWN,
	"<pagedown>":                   KEY_PAGEDOWN,
	"KEY_INSERT":                   KEY_INSERT,
	"<delete>":                     KEY_DELETE,
	"<del>":                        KEY_DELETE,
	"KEY_MACRO":                    KEY_MACRO,
	"KEY_MUTE":                     KEY_MUTE,
	"KEY_VOLUMEDOWN":               KEY_VOLUMEDOWN,
	"KEY_VOLUMEUP":                 KEY_VOLUMEUP,
	"KEY_POWER":                    KEY_POWER,
	"KEY_KPEQUAL":                  KEY_KPEQUAL,
	"KEY_KPPLUSMINUS":              KEY_KPPLUSMINUS,
	"KEY_PAUSE":                    KEY_PAUSE,
	"KEY_SCALE":                    KEY_SCALE,
	"KEY_KPCOMMA":                  KEY_KPCOMMA,
	"KEY_HANGEUL":                  KEY_HANGEUL,
	"KEY_HANGUEL":                  KEY_HANGUEL,
	"KEY_HANJA":                    KEY_HANJA,
	"KEY_YEN":                      KEY_YEN,
	"KEY_LEFTMETA":                 KEY_LEFTMETA,
	"KEY_RIGHTMETA":                KEY_RIGHTMETA,
	"KEY_COMPOSE":                  KEY_COMPOSE,
	"KEY_STOP":                     KEY_STOP,
	"KEY_AGAIN":                    KEY_AGAIN,
	"KEY_PROPS":                    KEY_PROPS,
	"KEY_UNDO":                     KEY_UNDO,
	"KEY_FRONT":                    KEY_FRONT,
	"KEY_COPY":                     KEY_COPY,
	"KEY_OPEN":                     KEY_OPEN,
	"KEY_PASTE":                    KEY_PASTE,
	"KEY_FIND":                     KEY_FIND,
	"KEY_CUT":                      KEY_CUT,
	"KEY_HELP":                     KEY_HELP,
	"KEY_MENU":                     KEY_MENU,
	"KEY_CALC":                     KEY_CALC,
	"KEY_SETUP":                    KEY_SETUP,
	"KEY_SLEEP":                    KEY_SLEEP,
	"KEY_WAKEUP":                   KEY_WAKEUP,
	"KEY_FILE":                     KEY_FILE,
	"KEY_SENDFILE":                 KEY_SENDFILE,
	"KEY_DELETEFILE":               KEY_DELETEFILE,
	"KEY_XFER":                     KEY_XFER,
	"KEY_PROG1":                    KEY_PROG1,
	"KEY_PROG2":                    KEY_PROG2,
	"KEY_WWW":                      KEY_WWW,
	"KEY_MSDOS":                    KEY_MSDOS,
	"KEY_COFFEE":                   KEY_COFFEE,
	"KEY_SCREENLOCK":               KEY_SCREENLOCK,
	"KEY_DIRECTION":                KEY_DIRECTION,
	"KEY_CYCLEWINDOWS":             KEY_CYCLEWINDOWS,
	"KEY_MAIL":                     KEY_MAIL,
	"KEY_BOOKMARKS":                KEY_BOOKMARKS,
	"KEY_COMPUTER":                 KEY_COMPUTER,
	"KEY_BACK":                     KEY_BACK,
	"KEY_FORWARD":                  KEY_FORWARD,
	"KEY_CLOSECD":                  KEY_CLOSECD,
	"KEY_EJECTCD":                  KEY_EJECTCD,
	"KEY_EJECTCLOSECD":             KEY_EJECTCLOSECD,
	"KEY_NEXTSONG":                 KEY_NEXTSONG,
	"KEY_PLAYPAUSE":                KEY_PLAYPAUSE,
	"KEY_PREVIOUSSONG":             KEY_PREVIOUSSONG,
	"KEY_STOPCD":                   KEY_STOPCD,
	"KEY_RECORD":                   KEY_RECORD,
	"KEY_REWIND":                   KEY_REWIND,
	"KEY_PHONE":                    KEY_PHONE,
	"KEY_ISO":                      KEY_ISO,
	"KEY_CONFIG":                   KEY_CONFIG,
	"KEY_HOMEPAGE":                 KEY_HOMEPAGE,
	"KEY_REFRESH":                  KEY_REFRESH,
	"KEY_EXIT":                     KEY_EXIT,
	"KEY_MOVE":                     KEY_MOVE,
	"KEY_EDIT":                     KEY_EDIT,
	"KEY_SCROLLUP":                 KEY_SCROLLUP,
	"KEY_SCROLLDOWN":               KEY_SCROLLDOWN,
	"KEY_KPLEFTPAREN":              KEY_KPLEFTPAREN,
	"KEY_KPRIGHTPAREN":             KEY_KPRIGHTPAREN,
	"KEY_NEW":                      KEY_NEW,
	"KEY_REDO":                     KEY_REDO,
	"KEY_F13":                      KEY_F13,
	"KEY_F14":                      KEY_F14,
	"KEY_F15":                      KEY_F15,
	"KEY_F16":                      KEY_F16,
	"KEY_F17":                      KEY_F17,
	"KEY_F18":                      KEY_F18,
	"KEY_F19":                      KEY_F19,
	"KEY_F20":                      KEY_F20,
	"KEY_F21":                      KEY_F21,
	"KEY_F22":                      KEY_F22,
	"KEY_F23":                      KEY_F23,
	"KEY_F24":                      KEY_F24,
	"KEY_PLAYCD":                   KEY_PLAYCD,
	"KEY_PAUSECD":                  KEY_PAUSECD,
	"KEY_PROG3":                    KEY_PROG3,
	"KEY_PROG4":                    KEY_PROG4,
	"KEY_DASHBOARD":                KEY_DASHBOARD,
	"KEY_SUSPEND":                  KEY_SUSPEND,
	"KEY_CLOSE":                    KEY_CLOSE,
	"KEY_PLAY":                     KEY_PLAY,
	"KEY_FASTFORWARD":              KEY_FASTFORWARD,
	"KEY_BASSBOOST":                KEY_BASSBOOST,
	"KEY_PRINT":                    KEY_PRINT,
	"KEY_HP":                       KEY_HP,
	"KEY_CAMERA":                   KEY_CAMERA,
	"KEY_SOUND":                    KEY_SOUND,
	"KEY_QUESTION":                 KEY_QUESTION,
	"KEY_EMAIL":                    KEY_EMAIL,
	"KEY_CHAT":                     KEY_CHAT,
	"KEY_SEARCH":                   KEY_SEARCH,
	"KEY_CONNECT":                  KEY_CONNECT,
	"KEY_FINANCE":                  KEY_FINANCE,
	"KEY_SPORT":                    KEY_SPORT,
	"KEY_SHOP":                     KEY_SHOP,
	"KEY_ALTERASE":                 KEY_ALTERASE,
	"KEY_CANCEL":                   KEY_CANCEL,
	"KEY_BRIGHTNESSDOWN":           KEY_BRIGHTNESSDOWN,
	"KEY_BRIGHTNESSUP":             KEY_BRIGHTNESSUP,
	"KEY_MEDIA":                    KEY_MEDIA,
	"KEY_SWITCHVIDEOMODE":          KEY_SWITCHVIDEOMODE,
	"KEY_KBDILLUMTOGGLE":           KEY_KBDILLUMTOGGLE,
	"KEY_KBDILLUMDOWN":             KEY_KBDILLUMDOWN,
	"KEY_KBDILLUMUP":               KEY_KBDILLUMUP,
	"KEY_SEND":                     KEY_SEND,
	"KEY_REPLY":                    KEY_REPLY,
	"KEY_FORWARDMAIL":              KEY_FORWARDMAIL,
	"KEY_SAVE":                     KEY_SAVE,
	"KEY_DOCUMENTS":                KEY_DOCUMENTS,
	"KEY_BATTERY":                  KEY_BATTERY,
	"KEY_BLUETOOTH":                KEY_BLUETOOTH,
	"KEY_WLAN":                     KEY_WLAN,
	"KEY_UWB":                      KEY_UWB,
	"KEY_UNKNOWN":                  KEY_UNKNOWN,
	"KEY_VIDEO_NEXT":               KEY_VIDEO_NEXT,
	"KEY_VIDEO_PREV":               KEY_VIDEO_PREV,
	"KEY_BRIGHTNESS_CYCLE":         KEY_BRIGHTNESS_CYCLE,
	"KEY_BRIGHTNESS_AUTO":          KEY_BRIGHTNESS_AUTO,
	"KEY_BRIGHTNESS_ZERO":          KEY_BRIGHTNESS_ZERO,
	"KEY_DISPLAY_OFF":              KEY_DISPLAY_OFF,
	"KEY_WWAN":                     KEY_WWAN,
	"KEY_WIMAX":                    KEY_WIMAX,
	"KEY_RFKILL":                   KEY_RFKILL,
	"KEY_MICMUTE":                  KEY_MICMUTE,
	"KEY_OK":                       KEY_OK,
	"KEY_SELECT":                   KEY_SELECT,
	"KEY_GOTO":                     KEY_GOTO,
	"KEY_CLEAR":                    KEY_CLEAR,
	"KEY_POWER2":                   KEY_POWER2,
	"KEY_OPTION":                   KEY_OPTION,
	"KEY_INFO":                     KEY_INFO,
	"KEY_TIME":                     KEY_TIME,
	"KEY_VENDOR":                   KEY_VENDOR,
	"KEY_ARCHIVE":                  KEY_ARCHIVE,
	"KEY_PROGRAM":                  KEY_PROGRAM,
	"KEY_CHANNEL":                  KEY_CHANNEL,
	"KEY_FAVORITES":                KEY_FAVORITES,
	"KEY_EPG":                      KEY_EPG,
	"KEY_PVR":                      KEY_PVR,
	"KEY_MHP":                      KEY_MHP,
	"KEY_LANGUAGE":                 KEY_LANGUAGE,
	"KEY_TITLE":                    KEY_TITLE,
	"KEY_SUBTITLE":                 KEY_SUBTITLE,
	"KEY_ANGLE":                    KEY_ANGLE,
	"KEY_ZOOM":                     KEY_ZOOM,
	"KEY_MODE":                     KEY_MODE,
	"KEY_KEYBOARD":                 KEY_KEYBOARD,
	"KEY_SCREEN":                   KEY_SCREEN,
	"KEY_PC":                       KEY_PC,
	"KEY_TV":                       KEY_TV,
	"KEY_TV2":                      KEY_TV2,
	"KEY_VCR":                      KEY_VCR,
	"KEY_VCR2":                     KEY_VCR2,
	"KEY_SAT":                      KEY_SAT,
	"KEY_SAT2":                     KEY_SAT2,
	"KEY_CD":                       KEY_CD,
	"KEY_TAPE":                     KEY_TAPE,
	"KEY_RADIO":                    KEY_RADIO,
	"KEY_TUNER":                    KEY_TUNER,
	"KEY_PLAYER":                   KEY_PLAYER,
	"KEY_TEXT":                     KEY_TEXT,
	"KEY_DVD":                      KEY_DVD,
	"KEY_AUX":                      KEY_AUX,
	"KEY_MP3":                      KEY_MP3,
	"KEY_AUDIO":                    KEY_AUDIO,
	"KEY_VIDEO":                    KEY_VIDEO,
	"KEY_DIRECTORY":                KEY_DIRECTORY,
	"KEY_LIST":                     KEY_LIST,
	"KEY_MEMO":                     KEY_MEMO,
	"KEY_CALENDAR":                 KEY_CALENDAR,
	"KEY_RED":                      KEY_RED,
	"KEY_GREEN":                    KEY_GREEN,
	"KEY_YELLOW":                   KEY_YELLOW,
	"KEY_BLUE":                     KEY_BLUE,
	"KEY_CHANNELUP":                KEY_CHANNELUP,
	"KEY_CHANNELDOWN":              KEY_CHANNELDOWN,
	"KEY_FIRST":                    KEY_FIRST,
	"KEY_LAST":                     KEY_LAST,
	"KEY_AB":                       KEY_AB,
	"KEY_NEXT":                     KEY_NEXT,
	"KEY_RESTART":                  KEY_RESTART,
	"KEY_SLOW":                     KEY_SLOW,
	"KEY_SHUFFLE":                  KEY_SHUFFLE,
	"KEY_BREAK":                    KEY_BREAK,
	"KEY_PREVIOUS":                 KEY_PREVIOUS,
	"KEY_DIGITS":                   KEY_DIGITS,
	"KEY_TEEN":                     KEY_TEEN,
	"KEY_TWEN":                     KEY_TWEN,
	"KEY_VIDEOPHONE":               KEY_VIDEOPHONE,
	"KEY_GAMES":                    KEY_GAMES,
	"KEY_ZOOMIN":                   KEY_ZOOMIN,
	"KEY_ZOOMOUT":                  KEY_ZOOMOUT,
	"KEY_ZOOMRESET":                KEY_ZOOMRESET,
	"KEY_WORDPROCESSOR":            KEY_WORDPROCESSOR,
	"KEY_EDITOR":                   KEY_EDITOR,
	"KEY_SPREADSHEET":              KEY_SPREADSHEET,
	"KEY_GRAPHICSEDITOR":           KEY_GRAPHICSEDITOR,
	"KEY_PRESENTATION":             KEY_PRESENTATION,
	"KEY_DATABASE":                 KEY_DATABASE,
	"KEY_NEWS":                     KEY_NEWS,
	"KEY_VOICEMAIL":                KEY_VOICEMAIL,
	"KEY_ADDRESSBOOK":              KEY_ADDRESSBOOK,
	"KEY_MESSENGER":                KEY_MESSENGER,
	"KEY_DISPLAYTOGGLE":            KEY_DISPLAYTOGGLE,
	"KEY_BRIGHTNESS_TOGGLE":        KEY_BRIGHTNESS_TOGGLE,
	"KEY_SPELLCHECK":               KEY_SPELLCHECK,
	"KEY_LOGOFF":                   KEY_LOGOFF,
	"KEY_DOLLAR":                   KEY_DOLLAR,
	"KEY_EURO":                     KEY_EURO,
	"KEY_FRAMEBACK":                KEY_FRAMEBACK,
	"KEY_FRAMEFORWARD":             KEY_FRAMEFORWARD,
	"KEY_CONTEXT_MENU":             KEY_CONTEXT_MENU,
	"KEY_MEDIA_REPEAT":             KEY_MEDIA_REPEAT,
	"KEY_10CHANNELSUP":             KEY_10CHANNELSUP,
	"KEY_10CHANNELSDOWN":           KEY_10CHANNELSDOWN,
	"KEY_IMAGES":                   KEY_IMAGES,
	"KEY_DEL_EOL":                  KEY_DEL_EOL,
	"KEY_DEL_EOS":                  KEY_DEL_EOS,
	"KEY_INS_LINE":                 KEY_INS_LINE,
	"KEY_DEL_LINE":                 KEY_DEL_LINE,
	"KEY_FN":                       KEY_FN,
	"KEY_FN_ESC":                   KEY_FN_ESC,
	"KEY_FN_F1":                    KEY_FN_F1,
	"KEY_FN_F2":                    KEY_FN_F2,
	"KEY_FN_F3":                    KEY_FN_F3,
	"KEY_FN_F4":                    KEY_FN_F4,
	"KEY_FN_F5":                    KEY_FN_F5,
	"KEY_FN_F6":                    KEY_FN_F6,
	"KEY_FN_F7":                    KEY_FN_F7,
	"KEY_FN_F8":                    KEY_FN_F8,
	"KEY_FN_F9":                    KEY_FN_F9,
	"KEY_FN_F10":                   KEY_FN_F10,
	"KEY_FN_F11":                   KEY_FN_F11,
	"KEY_FN_F12":                   KEY_FN_F12,
	"KEY_FN_1":                     KEY_FN_1,
	"KEY_FN_2":                     KEY_FN_2,
	"KEY_FN_D":                     KEY_FN_D,
	"KEY_FN_E":                     KEY_FN_E,
	"KEY_FN_F":                     KEY_FN_F,
	"KEY_FN_S":                     KEY_FN_S,
	"KEY_FN_B":                     KEY_FN_B,
	"KEY_BRL_DOT1":                 KEY_BRL_DOT1,
	"KEY_BRL_DOT2":                 KEY_BRL_DOT2,
	"KEY_BRL_DOT3":                 KEY_BRL_DOT3,
	"KEY_BRL_DOT4":                 KEY_BRL_DOT4,
	"KEY_BRL_DOT5":                 KEY_BRL_DOT5,
	"KEY_BRL_DOT6":                 KEY_BRL_DOT6,
	"KEY_BRL_DOT7":                 KEY_BRL_DOT7,
	"KEY_BRL_DOT8":                 KEY_BRL_DOT8,
	"KEY_BRL_DOT9":                 KEY_BRL_DOT9,
	"KEY_BRL_DOT10":                KEY_BRL_DOT10,
	"KEY_NUMERIC_0":                KEY_NUMERIC_0,
	"KEY_NUMERIC_1":                KEY_NUMERIC_1,
	"KEY_NUMERIC_2":                KEY_NUMERIC_2,
	"KEY_NUMERIC_3":                KEY_NUMERIC_3,
	"KEY_NUMERIC_4":                KEY_NUMERIC_4,
	"KEY_NUMERIC_5":                KEY_NUMERIC_5,
	"KEY_NUMERIC_6":                KEY_NUMERIC_6,
	"KEY_NUMERIC_7":                KEY_NUMERIC_7,
	"KEY_NUMERIC_8":                KEY_NUMERIC_8,
	"KEY_NUMERIC_9":                KEY_NUMERIC_9,
	"KEY_NUMERIC_STAR":             KEY_NUMERIC_STAR,
	"KEY_NUMERIC_POUND":            KEY_NUMERIC_POUND,
	"KEY_CAMERA_FOCUS":             KEY_CAMERA_FOCUS,
	"KEY_WPS_BUTTON":               KEY_WPS_BUTTON,
	"KEY_TOUCHPAD_TOGGLE":          KEY_TOUCHPAD_TOGGLE,
	"KEY_TOUCHPAD_ON":              KEY_TOUCHPAD_ON,
	"KEY_TOUCHPAD_OFF":             KEY_TOUCHPAD_OFF,
	"KEY_CAMERA_ZOOMIN":            KEY_CAMERA_ZOOMIN,
	"KEY_CAMERA_ZOOMOUT":           KEY_CAMERA_ZOOMOUT,
	"KEY_CAMERA_UP":                KEY_CAMERA_UP,
	"KEY_CAMERA_DOWN":              KEY_CAMERA_DOWN,
	"KEY_CAMERA_LEFT":              KEY_CAMERA_LEFT,
	"KEY_CAMERA_RIGHT":             KEY_CAMERA_RIGHT,
	"KEY_ATTENDANT_ON":             KEY_ATTENDANT_ON,
	"KEY_ATTENDANT_OFF":            KEY_ATTENDANT_OFF,
	"KEY_ATTENDANT_TOGGLE":         KEY_ATTENDANT_TOGGLE,
	"KEY_LIGHTS_TOGGLE":            KEY_LIGHTS_TOGGLE,
	"KEY_ALS_TOGGLE":               KEY_ALS_TOGGLE,
	"KEY_BUTTONCONFIG":             KEY_BUTTONCONFIG,
	"KEY_TASKMANAGER":              KEY_TASKMANAGER,
	"KEY_JOURNAL":                  KEY_JOURNAL,
	"KEY_CONTROLPANEL":             KEY_CONTROLPANEL,
	"KEY_APPSELECT":                KEY_APPSELECT,
	"KEY_SCREENSAVER":              KEY_SCREENSAVER,
	"KEY_VOICECOMMAND":             KEY_VOICECOMMAND,
	"KEY_BRIGHTNESS_MIN":           KEY_BRIGHTNESS_MIN,
	"KEY_BRIGHTNESS_MAX":           KEY_BRIGHTNESS_MAX,
	"KEY_KBDINPUTASSIST_PREV":      KEY_KBDINPUTASSIST_PREV,
	"KEY_KBDINPUTASSIST_NEXT":      KEY_KBDINPUTASSIST_NEXT,
	"KEY_KBDINPUTASSIST_PREVGROUP": KEY_KBDINPUTASSIST_PREVGROUP,
	"KEY_KBDINPUTASSIST_NEXTGROUP": KEY_KBDINPUTASSIST_NEXTGROUP,
	"KEY_KBDINPUTASSIST_ACCEPT":    KEY_KBDINPUTASSIST_ACCEPT,
	"KEY_KBDINPUTASSIST_CANCEL":    KEY_KBDINPUTASSIST_CANCEL,
	"KEY_MIN_INTERESTING":          KEY_MIN_INTERESTING,
	"KEY_MAX":                      KEY_MAX,
	"KEY_CNT":                      KEY_CNT,
}
