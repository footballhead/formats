package config_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sanctuary/formats/image/cel/config"
)

func TestConfs(t *testing.T) {
	var relCelPaths []string
	for relCelPath := range config.Confs {
		relCelPaths = append(relCelPaths, relCelPath)
	}
	sort.Strings(relCelPaths)

	for _, relCelPath := range relCelPaths {
		npixels, ok := npixelsMapping[relCelPath]
		if !ok {
			t.Errorf("unable to locate pixel count for %q", relCelPath)
			continue
		}

		// Get config and frame numbers with specific image dimensions.
		conf := config.Confs[relCelPath]

		// TODO: Check if this test is redundant, and may therefore be removed.

		// Verify pixel count for each frame with specific image dimensions.
		m := make(map[int]bool)
		for frameNum := range conf.FrameWidth {
			m[frameNum] = true
		}
		for frameNum := range conf.FrameHeight {
			m[frameNum] = true
		}
		var frameNums []int
		for frameNum := range m {
			frameNums = append(frameNums, frameNum)
		}
		sort.Ints(frameNums)
		for _, frameNum := range frameNums {
			width, ok := conf.FrameWidth[frameNum]
			if !ok {
				width = conf.W
			}
			height, ok := conf.FrameHeight[frameNum]
			if !ok {
				height = conf.H
			}
			got := width * height
			want := npixels[frameNum]
			if got != want {
				t.Errorf("%q: pixel count mismatch for frame number %d; expected %d, got %d", relCelPath, frameNum, want, got)
				continue
			}
		}

		// Verify pixel count for default dimension frames.
		for frameNum, want := range npixels {
			width, ok := conf.FrameWidth[frameNum]
			if !ok {
				width = conf.W
			}
			height, ok := conf.FrameHeight[frameNum]
			if !ok {
				height = conf.H
			}
			got := width * height
			if got != want {
				t.Errorf("%q: pixel count mismatch for frame number %d; expected %d, got %d", relCelPath, frameNum, want, got)
				continue
			} else {
				fmt.Println("PASS:", relCelPath)
			}
		}
	}
}

// TODO: Add once supported.
//
// * "monsters/unrav/unravw.cel"

// TODO: Add once support for CEL archives has been added.
//
// * "towners/animals/cow.cel"
// * "towners/smith/smithw.cel"
// * "towners/townwmn1/wmnw.cel"
// * "towners/twnf/twnfw.cel"

// npixelsMapping maps from CEL file name to pixel count. The pixel count slice
// maps from frame number to pixel count of the given frame. If each frame of
// the CEL has the same pixel count, then the pixel count slice contains a
// single element.
var npixelsMapping = map[string][]int{
	"ctrlpan/golddrop.cel":          {35496},
	"ctrlpan/p8bulbs.cel":           {7744},
	"ctrlpan/p8but2.cel":            {1056},
	"ctrlpan/panel8.cel":            {92160},
	"ctrlpan/panel8bu.cel":          {1349},
	"ctrlpan/smaltext.cel":          {143},
	"ctrlpan/spelicon.cel":          {3136},
	"ctrlpan/talkbutt.cel":          {976},
	"ctrlpan/talkpanl.cel":          {92160},
	"data/bigtgold.cel":             {2070},
	"data/char.cel":                 {112640},
	"data/charbut.cel":              {2090, 902, 902, 902, 902, 902, 902, 902, 902},
	"data/diabsmal.cel":             {29600},
	"data/inv/inv.cel":              {112640},
	"data/inv/inv_rog.cel":          {112640},
	"data/inv/inv_sor.cel":          {112640},
	"data/inv/objcurs.cel":          {946, 1021, 1002, 1004, 1002, 1013, 1005, 1007, 1015, 1002, 808, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 766, 792, 766, 766, 766, 1550, 1550, 1550, 1568, 1550, 1550, 2334, 2334, 2334, 2334, 2334, 2334, 2334, 2334, 2334, 2334, 2334, 2334, 2334, 2334, 2334, 2334, 2334, 2334, 2334, 3090, 3090, 3090, 3090, 3128, 3090, 3090, 3090, 3090, 3090, 3090, 3090, 3090, 3090, 3090, 3090, 3090, 3090, 3090, 3090, 3090, 3090, 3090, 3132, 3090, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4709, 4709, 4677, 4658, 4658, 4663, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4661, 4670, 4658, 4658, 4701, 4697, 4658, 4658, 4658, 4658, 4686, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4703, 4658, 4658, 4658, 4658, 4658, 4658, 4658, 4689, 4668, 4658, 4665, 4708, 4672, 4658, 4666, 4658},
	"data/medtexts.cel":             {484},
	"data/optbar.cel":               {9184},
	"data/option.cel":               {756},
	"data/pentspin.cel":             {2304},
	"data/pentspn2.cel":             {144},
	"data/quest.cel":                {112640},
	"data/spellbk.cel":              {112640},
	"data/spellbkb.cel":             {2204},
	"data/spelli2.cel":              {1406},
	"data/square.cel":               {8172},
	"data/textbox.cel":              {179073},
	"data/textbox2.cel":             {82113},
	"data/textslid.cel":             {144},
	"gendata/cut2.cel":              {307200},
	"gendata/cut3.cel":              {307200},
	"gendata/cut4.cel":              {307200},
	"gendata/cutgate.cel":           {307200},
	"gendata/cutl1d.cel":            {307200},
	"gendata/cutportl.cel":          {307200},
	"gendata/cutportr.cel":          {307200},
	"gendata/cutstart.cel":          {307200},
	"gendata/cuttt.cel":             {307200},
	"gendata/quotes.cel":            {307200},
	"items/armor2.cel":              {15274},
	"items/axe.cel":                 {15274},
	"items/axeflip.cel":             {15274},
	"items/bldstn.cel":              {15274},
	"items/bottle.cel":              {15274},
	"items/bow.cel":                 {15274},
	"items/cleaver.cel":             {15274},
	"items/crownf.cel":              {12202},
	"items/duricons.cel":            {1024},
	"items/fanvil.cel":              {15274},
	"items/fbook.cel":               {15274},
	"items/fbow.cel":                {15274},
	"items/fbrain.cel":              {15274},
	"items/fbttle.cel":              {15274},
	"items/fbttlebb.cel":            {15274},
	"items/fbttlebl.cel":            {15274},
	"items/fbttlebr.cel":            {15274},
	"items/fbttleby.cel":            {15274},
	"items/fbttledb.cel":            {15274},
	"items/fbttledy.cel":            {15274},
	"items/fbttleor.cel":            {15274},
	"items/fbttlewh.cel":            {15274},
	"items/fear.cel":                {12202},
	"items/feye.cel":                {15274},
	"items/fheart.cel":              {15274},
	"items/flazstaf.cel":            {15274},
	"items/fmush.cel":               {15274},
	"items/food.cel":                {12202},
	"items/fplatear.cel":            {15274},
	"items/goldflip.cel":            {15274},
	"items/helmut.cel":              {15274},
	"items/innsign.cel":             {15274},
	"items/larmor.cel":              {12202},
	"items/mace.cel":                {15274},
	"items/manaflip.cel":            {15274},
	"items/map/mapz0000.cel":        {225280},
	"items/map/mapz0001.cel":        {225280},
	"items/map/mapz0002.cel":        {225280},
	"items/map/mapz0003.cel":        {225280},
	"items/map/mapz0004.cel":        {225280},
	"items/map/mapz0005.cel":        {225280},
	"items/map/mapz0006.cel":        {225280},
	"items/map/mapz0007.cel":        {225280},
	"items/map/mapz0008.cel":        {225280},
	"items/map/mapz0009.cel":        {225280},
	"items/map/mapz0010.cel":        {225280},
	"items/map/mapz0011.cel":        {225280},
	"items/map/mapz0012.cel":        {225280},
	"items/map/mapz0013.cel":        {225280},
	"items/map/mapz0014.cel":        {225280},
	"items/map/mapz0015.cel":        {225280},
	"items/map/mapz0016.cel":        {225280},
	"items/map/mapz0017.cel":        {225280},
	"items/map/mapz0018.cel":        {225280},
	"items/map/mapz0019.cel":        {225280},
	"items/map/mapz0020.cel":        {225280},
	"items/map/mapz0021.cel":        {225280},
	"items/map/mapz0022.cel":        {225280},
	"items/map/mapz0023.cel":        {225280},
	"items/map/mapz0024.cel":        {225280},
	"items/map/mapz0025.cel":        {225280},
	"items/map/mapz0026.cel":        {225280},
	"items/map/mapz0027.cel":        {225280},
	"items/map/mapz0028.cel":        {225280},
	"items/map/mapz0029.cel":        {225280},
	"items/map/mapz0030.cel":        {225280},
	"items/map/mapzdoom.cel":        {225280},
	"items/ring.cel":                {15274},
	"items/rock.cel":                {9130},
	"items/scroll.cel":              {15274},
	"items/shield.cel":              {15274},
	"items/staff.cel":               {15274},
	"items/swrdflip.cel":            {15274},
	"items/wand.cel":                {15274},
	"items/wshield.cel":             {12202},
	"levels/l1data/l1s.cel":         {10186},
	"levels/l2data/l2s.cel":         {10186},
	"levels/towndata/towns.cel":     {14336},
	"missiles/flamel1.cel":          {12170},
	"missiles/flamel10.cel":         {9130},
	"missiles/flamel11.cel":         {12170},
	"missiles/flamel12.cel":         {9130},
	"missiles/flamel13.cel":         {12170},
	"missiles/flamel14.cel":         {9130},
	"missiles/flamel15.cel":         {12170},
	"missiles/flamel16.cel":         {9130},
	"missiles/flamel2.cel":          {9130},
	"missiles/flamel3.cel":          {12170},
	"missiles/flamel4.cel":          {9130},
	"missiles/flamel5.cel":          {12170},
	"missiles/flamel6.cel":          {9130},
	"missiles/flamel7.cel":          {12170},
	"missiles/flamel8.cel":          {9130},
	"missiles/flamel9.cel":          {12170},
	"missiles/flames1.cel":          {12170},
	"missiles/flames10.cel":         {9130},
	"missiles/flames11.cel":         {12170},
	"missiles/flames12.cel":         {9130},
	"missiles/flames13.cel":         {12170},
	"missiles/flames14.cel":         {9130},
	"missiles/flames15.cel":         {12170},
	"missiles/flames16.cel":         {9130},
	"missiles/flames2.cel":          {9130},
	"missiles/flames3.cel":          {12170},
	"missiles/flames4.cel":          {9130},
	"missiles/flames5.cel":          {12170},
	"missiles/flames6.cel":          {9130},
	"missiles/flames7.cel":          {12170},
	"missiles/flames8.cel":          {9130},
	"missiles/flames9.cel":          {12170},
	"missiles/flaml1.cel":           {16266},
	"missiles/flaml2.cel":           {16266},
	"missiles/flaml3.cel":           {16266},
	"missiles/flaml4.cel":           {16266},
	"missiles/flaml5.cel":           {16266},
	"missiles/flaml6.cel":           {16266},
	"missiles/flaml7.cel":           {16266},
	"missiles/flaml8.cel":           {16266},
	"missiles/flams1.cel":           {16266},
	"missiles/flams2.cel":           {16266},
	"missiles/flams3.cel":           {16266},
	"missiles/flams4.cel":           {16266},
	"missiles/flams5.cel":           {16266},
	"missiles/flams6.cel":           {16266},
	"missiles/flams7.cel":           {16266},
	"missiles/flams8.cel":           {16266},
	"missiles/mindmace.cel":         {9130},
	"missiles/sentfr.cel":           {9130, 9187, 9130},
	"missiles/sentout.cel":          {9130},
	"missiles/sentup.cel":           {9130},
	"monsters/acid/acidpud.cel":     {12170},
	"monsters/magma/magball1.cel":   {16266},
	"monsters/magma/magball2.cel":   {16266},
	"monsters/magma/magball3.cel":   {16266},
	"monsters/magma/magball4.cel":   {16266},
	"monsters/magma/magball5.cel":   {16266},
	"monsters/magma/magball6.cel":   {16266},
	"monsters/magma/magball7.cel":   {16266},
	"monsters/magma/magball8.cel":   {16266},
	"monsters/magma/magblos.cel":    {16266},
	"monsters/rhino/rhinos1.cel":    {20362},
	"monsters/rhino/rhinos2.cel":    {20362},
	"monsters/rhino/rhinos3.cel":    {20362},
	"monsters/rhino/rhinos4.cel":    {20362, 20400, 20362, 20362, 20362, 20362},
	"monsters/rhino/rhinos5.cel":    {20362, 20362, 20419, 20362, 20404, 20411},
	"monsters/rhino/rhinos6.cel":    {20362, 20362, 20362, 20362, 20423, 20362},
	"monsters/rhino/rhinos7.cel":    {20362},
	"monsters/rhino/rhinos8.cel":    {20362},
	"monsters/succ/flare.cel":       {16266},
	"monsters/succ/flarexp.cel":     {16266, 16266, 16266, 16363, 16266, 16367, 16266},
	"monsters/thin/lghning.cel":     {9130},
	"objects/altboy.cel":            {16266},
	"objects/angel.cel":             {12202},
	"objects/armstand.cel":          {9130},
	"objects/banner.cel":            {9130},
	"objects/barrel.cel":            {9130, 9130, 9130, 9130, 9130, 9130, 9130, 9130, 9198},
	"objects/barrelex.cel":          {9130, 9130, 9130, 9130, 9130, 9130, 9130, 9130, 9198, 9198},
	"objects/bcase.cel":             {12202},
	"objects/bkslbrnt.cel":          {9130},
	"objects/bkurns.cel":            {9130, 9130, 9130, 9130, 9130, 9184, 9130, 9130, 9130, 9130},
	"objects/bloodfnt.cel":          {9130},
	"objects/book1.cel":             {9130},
	"objects/book2.cel":             {9130},
	"objects/bshelf.cel":            {12202},
	"objects/burncros.cel":          {25482},
	"objects/candlabr.cel":          {9130},
	"objects/candle.cel":            {9130},
	"objects/candle2.cel":           {9130},
	"objects/cauldren.cel":          {9130},
	"objects/chest1.cel":            {9130},
	"objects/chest2.cel":            {9130},
	"objects/chest3.cel":            {9130},
	"objects/cruxsk1.cel":           {12202, 12202, 12202, 12202, 12202, 12202, 12202, 12202, 12202, 12202, 12202, 12202, 12263, 12263, 12263},
	"objects/cruxsk2.cel":           {12202, 12202, 12202, 12202, 12202, 12202, 12202, 12202, 12202, 12202, 12202, 12202, 12263, 12263, 12263},
	"objects/cruxsk3.cel":           {12202, 12202, 12202, 12202, 12202, 12202, 12263, 12202, 12202, 12202, 12202, 12202, 12263, 12263, 12263},
	"objects/decap.cel":             {9130},
	"objects/dirtfall.cel":          {15274},
	"objects/explod1.cel":           {16266},
	"objects/explod2.cel":           {16266},
	"objects/firewal1.cel":          {20362},
	"objects/flame1.cel":            {9130},
	"objects/flame3.cel":            {12202},
	"objects/ghost.cel":             {16266},
	"objects/goatshrn.cel":          {9130},
	"objects/l1braz.cel":            {10186},
	"objects/l1doors.cel":           {10186},
	"objects/l2doors.cel":           {8138},
	"objects/l3doors.cel":           {8138},
	"objects/lever.cel":             {9130},
	"objects/lshrineg.cel":          {16266},
	"objects/lzstand.cel":           {16266},
	"objects/mcirl.cel":             {9130},
	"objects/mfountn.cel":           {16266},
	"objects/miniwatr.cel":          {8138},
	"objects/mushptch.cel":          {9130},
	"objects/nude2.cel":             {16266},
	"objects/pedistl.cel":           {9130},
	"objects/pfountn.cel":           {16266},
	"objects/prsrplt1.cel":          {9179, 9180, 9180, 9179, 9130, 9130, 9130, 9130, 9180, 9180},
	"objects/rockstan.cel":          {9130},
	"objects/rshrineg.cel":          {16266},
	"objects/sarc.cel":              {12170},
	"objects/skulfire.cel":          {9130},
	"objects/skulpile.cel":          {9130},
	"objects/skulstik.cel":          {9130},
	"objects/switch2.cel":           {9130},
	"objects/switch3.cel":           {9130},
	"objects/switch4.cel":           {9130},
	"objects/tfountn.cel":           {12170},
	"objects/tnudem.cel":            {16266},
	"objects/tnudew.cel":            {16266},
	"objects/traphole.cel":          {9162},
	"objects/tsoul.cel":             {12170, 12170, 12170, 12254, 12170, 12235},
	"objects/vapor1.cel":            {16266, 16266, 16299, 16266, 16266, 16266, 16366, 16266, 16266, 16266, 16266, 16266, 16266},
	"objects/water.cel":             {20362},
	"objects/waterjug.cel":          {9130},
	"objects/weapstnd.cel":          {9130},
	"objects/wtorch1.cel":           {12202},
	"objects/wtorch2.cel":           {12202},
	"objects/wtorch3.cel":           {12202},
	"objects/wtorch4.cel":           {12202},
	"towners/butch/deadguy.cel":     {9130},
	"towners/drunk/twndrunk.cel":    {9130},
	"towners/healer/healer.cel":     {9130},
	"towners/priest/priest8.cel":    {9130},
	"towners/smith/smithn.cel":      {9130},
	"towners/strytell/strytell.cel": {9130},
	"towners/townboy/pegkid1.cel":   {6058},
	"towners/townwmn1/witch.cel":    {9130},
	"towners/townwmn1/wmnn.cel":     {9130},
	"towners/twnf/twnfn.cel":        {9130},
}
