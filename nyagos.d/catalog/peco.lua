nyagos.alias.peco_temp_out = function()
    for _,val in ipairs(share.peco_temp_out) do
        nyagos.write(val,"\n")
    end
end

nyagos.bindkey("C-o",function(this)
    local word,pos = this:lastword()
    share.peco_temp_out = nyagos.glob(word.."*")
    local result=nyagos.eval('peco_temp_out | peco')
    this:call("CLEAR_SCREEN")
    if string.find(result," ",1,true) then
        result = '"'..result..'"'
    end
    assert( this:replacefrom(pos,result) )
end)

nyagos.bindkey("C_R", function(this)
    local path = nyagos.pathjoin(nyagos.env.appdata,'NYAOS_ORG\\nyagos.history')
    local result = nyagos.eval('peco < '..path)
    this:call("CLEAR_SCREEN")
    return result
end)

nyagos.bindkey("M_H" , function(this)
    local result = nyagos.eval('cd --history | peco')
    this:call("CLEAR_SCREEN")
    if string.find(result,' ') then
        result = '"'..result..'"'
    end
    return result
end)